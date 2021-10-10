package main

import (
	"context"
	"fmt"
	pb "golang.conradwood.net/apis/logservice"
	"sync"
)

const (
	MAX_LOG_BUFS = 100
)

var (
	logidctr = uint64(0)
	ll       sync.Mutex
	logs     []*logbuf
)

type logbuf struct {
	logid uint64
	peer  string
	log   *pb.LogRequest
}

func addToBuf(ctx context.Context, peer string, lr *pb.LogRequest) {
	lb := &logbuf{peer: peer, log: lr}
	ll.Lock()
	logidctr++
	lb.logid = logidctr
	if len(logs) >= MAX_LOG_BUFS {
		logs = append(logs[1:], lb)

	} else {
		logs = append(logs, lb)
	}
	ll.Unlock()
}

/***************************************************************************************
******** BIG FAT WARNING    ----- READ ME --------
******** BIG FAT WARNING    ----- READ ME --------

* here's a funny one:
* if you print to stdout here, then every time a client will tail -f our logs
* then it'll be an endless loop of following the output for this function
* basically, tail -f calls this function, so don't output to stdout

******** BIG FAT WARNING    ----- READ ME --------
******** BIG FAT WARNING    ----- READ ME --------
***************************************************************************************/
func (s *LogService) GetLogCommandStdout(ctx context.Context, lr *pb.GetLogRequest) (*pb.GetLogResponse, error) {
	res := &pb.GetLogResponse{}
	for _, l := range logs {
		if int64(l.logid) < lr.MinimumLogID {
			continue
		}
		res.Entries = append(res.Entries, l.LogEntry()...)
	}
	if *debug {
		fmt.Printf("Returning %d log entries\n", len(res.Entries))
	}
	return res, nil
}
func (l *logbuf) LogEntry() []*pb.LogEntry {
	var res []*pb.LogEntry
	for _, line := range l.log.Lines {
		le := &pb.LogEntry{
			ID:       l.logid,
			Host:     l.peer,
			UserName: "",
			Occured:  uint64(line.Time),
			AppDef:   l.log.AppDef,
			Line:     line.Line,
			Status:   line.Status,
		}
		res = append(res, le)
	}
	return res
}
