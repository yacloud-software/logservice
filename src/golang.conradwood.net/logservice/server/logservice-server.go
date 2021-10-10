package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"golang.conradwood.net/go-easyops/prometheus"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/logservice"
	"golang.conradwood.net/go-easyops/cmdline"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/stack"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// static variables for flag parser
var (
	logfile          *os.File
	logfileName      = flag.String("logfile", "/var/log/logservice/full.log", "logfile to write to")
	port             = flag.Int("port", 10000, "The server port")
	debug            = flag.Bool("debug", false, "turn debug output on - DANGEROUS DO NOT USE IN PRODUCTION!")
	clean_on_startup = flag.Bool("clean_on_startup", false, "if true, removes old log files from the database on startup")
	reqCounter       = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "logservice_requests",
			Help: "requests to log stuff received",
		},
		[]string{"appname", "repositoryid"},
	)
	lineCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "logservice_lines",
			Help: "number of lines logged",
		},
		[]string{"appname", "repositoryid"},
	)
	byteCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "logservice_bytes",
			Help: "number of bytes logged",
		},
		[]string{"appname", "repositoryid"},
	)
	failCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "logservice_failed_requests",
			Help: "requests to log stuff received",
		},
		[]string{"appname", "repositoryid"},
	)
)

type pqDB struct {
	once    sync.Once // guards init of running
	running bool      // whether port 5432 is listening
}

func (p *pqDB) Running() bool {
	p.once.Do(func() {
		c, err := net.Dial("tcp", "localhost:5432")
		if err == nil {
			p.running = true
			c.Close()
		}
	})
	return p.running
}

// callback from the compound initialisation
func st(server *grpc.Server) error {
	s := new(LogService)
	// Register the handler object
	pb.RegisterLogServiceServer(server, s)
	return nil
}

func main() {
	var err error
	flag.Parse() // parse stuff. see "var" section above
	prometheus.MustRegister(reqCounter, failCounter, lineCounter, byteCounter)
	go rotate_loop()
	sd := server.NewServerDef()
	sd.Port = *port
	sd.Register = st
	sd.NoAuth = true
	err = server.ServerStartup(sd)
	if err != nil {
		fmt.Printf("failed to start server: %s\n", err)
	}
	fmt.Printf("Done\n")
	return

}

/**********************************
* implementing the functions here:
***********************************/
type LogService struct{}

/***************************************************************************************
******** BIG FAT WARNING    ----- READ ME --------
******** BIG FAT WARNING    ----- READ ME --------

* here's a funny one:
* if you print to stdout here, then it will be echoed back to you
* creating an endless loop.
* that's because we are also running in a service that logs
* stdout to us

******** BIG FAT WARNING    ----- READ ME --------
******** BIG FAT WARNING    ----- READ ME --------
***************************************************************************************/
func (s *LogService) LogCommandStdout(ctx context.Context, lr *pb.LogRequest) (*pb.LogResponse, error) {
	peer, ok := peer.FromContext(ctx)
	if !ok {
		return nil, errors.New("Error getting peer ")
	}
	peerhost, _, err := net.SplitHostPort(peer.Addr.String())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid peer: %v", peer))
	}
	addToBuf(ctx, peerhost, lr)
	rotate()
	l := prometheus.Labels{
		"appname":      lr.AppDef.Appname,
		"repositoryid": fmt.Sprintf("%d", lr.AppDef.RepoID),
	}
	reqCounter.With(l).Inc()
	lineCounter.With(l).Add(float64(len(lr.Lines)))
	bc := 0
	for _, l := range lr.Lines {
		bc = bc + len(l.Line)
	}
	byteCounter.With(l).Add(float64(bc))

	if *debug {
		fmt.Printf("Logging %d lines\n", len(lr.Lines))
	}
	if logfile == nil {
		logfile, err = utils.OpenWriteFile(*logfileName)
		if err != nil {
			fmt.Printf("Failed to open file: %s\n", err)
		}
	}
	appname := filepath.Base(lr.AppDef.Appname)
	for _, ll := range lr.Lines {
		line := ll.Line
		if len(line) > 999 {
			line = line[0:999]
		}
		ts := time.Now().Format("2/1/2006 15:04:05.000")
		sline := fmt.Sprintf("[%s] [%s] [%s]: \"%s\"\n", ts, peerhost, appname, line)
		if !cmdline.Datacenter() {
			fmt.Print(sline)
		}
		if logfile != nil {
			logfile.WriteString(sline)
			stack.Get(stackName(lr.AppDef)).Add(sline)
		}
	}
	resp := pb.LogResponse{}
	return &resp, nil
}
func stackName(l *pb.LogAppDef) string {
	return fmt.Sprintf("%s_%d_%s_%s", l.Appname, l.RepoID, l.Groupname, l.Namespace)
}

/***************************************************************************************
* retrieve applications final words ;)
* it's a specially useful case to get the logoutput from an application just before
* it terminated
***************************************************************************************/
func (s *LogService) GetAppLastEntries(ctx context.Context, lr *pb.GetHostLogRequest) (*pb.GetHostLogResponse, error) {
	return &pb.GetHostLogResponse{}, nil
}
func (s *LogService) GetApps(ctx context.Context, req *common.Void) (*pb.GetAppsResponse, error) {
	return &pb.GetAppsResponse{}, nil
}
func (s *LogService) CloseLog(ctx context.Context, req *pb.CloseLogRequest) (*common.Void, error) {
	var err error
	peer, ok := peer.FromContext(ctx)
	if !ok {
		return nil, errors.New("Error getting peer ")
	}
	peerhost, _, err := net.SplitHostPort(peer.Addr.String())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid peer: %v", peer))
	}
	rotate()
	if logfile == nil {
		logfile, err = utils.OpenWriteFile(*logfileName)
		if err != nil {
			fmt.Printf("Failed to open file: %s\n", err)
		}
	}
	a := req.AppDef
	appname := filepath.Base(a.Appname)
	line := fmt.Sprintf("==== Close log: exit code %d for %s/%s/%s ======= ", req.ExitCode, a.Namespace, a.Groupname, a.Appname)
	if len(line) > 999 {
		line = line[0:999]
	}
	ts := time.Now().Format("2/1/2006 15:04:05.000")
	sline := fmt.Sprintf("[%s] [%s] [%s]: \"%s\"\n", ts, peerhost, appname, line)
	if !cmdline.Datacenter() {
		fmt.Print(sline)
	}
	if logfile != nil {
		logfile.WriteString(sline)
	}
	ls := stack.Get(stackName(req.AppDef))
	go checkPanic(req.AppDef, ls.Get())
	return &common.Void{}, nil
}
