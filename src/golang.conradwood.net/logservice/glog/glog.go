package main

// the "GuruLog" client (glog)
// there's the logservice-client as well,
// but it's more like a testing tool than a userfriendly "log"
// we'll try with this tool to provide a more intuitive command line

import (
    "golang.conradwood.net/go-easyops/authremote"
	//"context"
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/logservice"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/utils"
	"os"
)

// static variables for flag parser
var (
	debug      = flag.Bool("debug", false, "debug this application")
	numentries = flag.Int("n", 100, "Amount of lines to retrieve")
	logServer  pb.LogServiceClient
	match      string
)

func main() {
	flag.Parse()
	logServer = pb.NewLogServiceClient(client.Connect("logservice.LogService"))
	filter := flag.Args()
	if len(filter) < 2 {
		flag.Usage()
		fmt.Printf("Usage: glog command [app1] [app2]\n")
		fmt.Printf("            command == [ls|tail|close]\n")
		fmt.Printf("             ls - list all applications\n")
		fmt.Printf("             tail - list most recent log entries (optional -f to follow)\n")
		fmt.Printf("             close - list most recent log entries just before application terminated\n")
		os.Exit(10)
	}
	com := filter[0]
	match = filter[1]
	if com == "close" {
		closedLog()
	} else {
		fmt.Printf("Invalid command: \"%s\"\n", com)
		os.Exit(10)
	}
}

func closedLog() {
	lr := &pb.GetLogRequest{}
	addFilters(lr)
	mli := int64(0 - *numentries)
	hr := &pb.GetHostLogRequest{MinimumLogID: mli}
	hr.LogFilter = lr.LogFilter
	res, err := logServer.GetAppLastEntries(authremote.Context(), hr)
	utils.Bail("Failed to get logs", err)
	for _, rp := range res.Entries {
		sortEntries(rp.Entries)
		for _, r := range rp.Entries {
			printLogEntry(r)
		}
	}
}
