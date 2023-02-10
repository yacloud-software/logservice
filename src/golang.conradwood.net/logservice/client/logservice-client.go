package main

import (
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/logservice"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/logger"
	"golang.conradwood.net/go-easyops/utils"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"time"
)

// static variables for flag parser
var (
	minappnamelen = 0
	match         = flag.String("filter", "", "fuzzy magic match filter to apply to logentries' application. Multiple matches delimited by comma, e.g. \"auth,registry\"")
	appName       = flag.String("appname", "", "The name of the application to log or to filter on")
	repo          = flag.Uint64("repository", 0, "The name of the repository to log")
	groupname     = flag.String("groupname", "", "The name of the group to log or to filter on")
	namespace     = flag.String("namespace", "", "the namespace to log or to filter on")
	deplid        = flag.String("deploymentid", "", "The deployment id to log")
	sid           = flag.String("startupid", "", "The startup id to log or to filter on")
	followFlag    = flag.Bool("f", false, "follow (tail -f like)")
	maxLines      = flag.Int("n", 500, "Maximum lines to retrieve")
	long          = flag.Bool("l", false, "long log entry output")
	logServer     pb.LogServiceClient
)

func main() {
	flag.Parse()
	logServer = pb.NewLogServiceClient(client.Connect("logservice.LogService"))
	lines := flag.Args()

	if *followFlag {
		follow()
		os.Exit(0)
	}
	if len(lines) == 0 {
		showLog()
		os.Exit(0)
	}
	queue, err := logger.NewAsyncLogQueue(*appName, *repo, 0, *groupname, *namespace, *deplid)
	utils.Bail("Failed to create log queue", err)
	for _, line := range lines {
		queue.LogCommandStdout(line, "EXECUSER")
		fmt.Printf("Logging: %s\n", line)
	}

	time.Sleep(5 * time.Second)
	err = queue.Flush()
	utils.Bail("Failed to send log", err)
	fmt.Printf("Done.\n")
}

func showLog() {

	ctx := authremote.Context()
	minlog := int64(0 - *maxLines)
	glr := pb.GetLogRequest{
		MinimumLogID: minlog,
	}
	addFilters(&glr)
	lr, err := logServer.GetLogCommandStdout(ctx, &glr)
	printApps(err)
	lastDate := ""
	for _, entry := range lr.Entries {
		printLogEntry(entry, &lastDate)
		if int64(entry.ID) >= minlog {
			minlog = int64(entry.ID)
		}
	}
	time.Sleep(1 * time.Second)
}

func follow() {

	minlog := int64(-20)
	i := 0
	lastDate := ""
	for {
		glr := pb.GetLogRequest{
			MinimumLogID: minlog,
		}
		addFilters(&glr)
		fmt.Fprintf(os.Stderr, "Querying %d (logid: %d)...\r", i, minlog)
		i++
		ctx := authremote.Context()
		lr, err := logServer.GetLogCommandStdout(ctx, &glr)
		if err == nil {
			sortEntries(lr.Entries)
			for _, entry := range lr.Entries {
				printLogEntry(entry, &lastDate)
				if int64(entry.ID) >= minlog {
					minlog = int64(entry.ID) + 1
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func sortEntries(entries []*pb.LogEntry) {
	sort.Slice(
		entries,
		func(j, k int) bool {
			return entries[j].ID < entries[k].ID
		},
	)
}

func printLogEntry(e *pb.LogEntry, lastDate *string) {
	t := time.Unix(int64(e.Occured), 0)
	bin := filepath.Base(e.AppDef.Appname)
	if *long {
		host := e.Host
		for len(host) < 15 {
			host = " " + host
		}
		ts := t.Format("2006-01-02 15:04:05.999999999 -0700 MST")
		fmt.Printf("%s %s %s repo:%03d app:%s: %s\n", ts, host, e.Status, e.AppDef.RepoID, bin, e.Line)
	} else {
		ts := t.Format("2006-01-02 15:04:05")
		if len(bin) > minappnamelen {
			minappnamelen = len(bin)
		}
		for len(bin) < minappnamelen {
			bin = " " + bin
		}
		fmt.Printf("%s %s: %s\n", ts, bin, e.Line)
	}
}

func addFilters(glr *pb.GetLogRequest) {
	if strings.Contains(*match, ",") {
		ads := filterToApps()
		for _, ad := range ads {
			lf := &pb.LogFilter{AppDef: ad}
			glr.LogFilter = append(glr.LogFilter, lf)
		}
		return
	}
	lf := &pb.LogFilter{
		FuzzyMatch: *match,
	}
	la := &pb.LogAppDef{
		Appname:   *appName,
		Groupname: *groupname,
		Namespace: *namespace,
		StartupID: *sid,
		RepoID:    *repo,
	}
	lf.AppDef = la
	glr.LogFilter = append(glr.LogFilter, lf)
}

// given a comma-delimeted string, will find logappdefs
func filterToApps() []*pb.LogAppDef {
	ctx := authremote.Context()
	x, xe := logServer.GetApps(ctx, &common.Void{})
	utils.Bail("Failed to get apps", xe)
	var res []*pb.LogAppDef
	for _, ld := range x.AppDef {
		fmt.Printf("Application: \"%s\", Namespace: \"%s\", Repository: \"%d\", Groupname: \"%s\"\n", ld.Appname, ld.Namespace, ld.RepoID, ld.Groupname)
		res = append(res, ld)
	}
	return res
}

func printApps(err error) {
	if err == nil {
		return
	}
	fmt.Printf("Getting available apps...\n")
	ctx := authremote.Context()
	x, xe := logServer.GetApps(ctx, &common.Void{})
	if xe == nil {
		for _, ld := range x.AppDef {
			fmt.Printf("Application: \"%s\", Namespace: \"%s\", Repository: \"%s\", Groupname: \"%s\"\n", ld.Appname, ld.Namespace, ld.Repository, ld.Groupname)
		}
	}
	utils.Bail("Failed to get Logcommandstdout", err)

}

func signalNotify(interrupt chan<- os.Signal) {
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
}
