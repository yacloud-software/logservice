package main

import (
    "golang.conradwood.net/go-easyops/authremote"
	"fmt"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/logservice"
	"golang.conradwood.net/go-easyops/utils"
	"sort"
	"strings"
	"time"
)

func sortEntries(entries []*pb.LogEntry) {
	sort.Slice(
		entries,
		func(j, k int) bool {
			return entries[j].ID < entries[k].ID
		},
	)
}

func printLogEntry(e *pb.LogEntry) {
	t := time.Unix(int64(e.Occured), 0)
	ts := t.String()

	fmt.Printf("%s %d %s %s repo:%s group:%s app:%s: %s\n", ts, e.ID, e.Host, e.Status,
		e.AppDef.Repository, e.AppDef.Groupname, e.AppDef.Appname,
		e.Line)
}

func addFilters(glr *pb.GetLogRequest) {
	ads := filterToApps()
	for _, ad := range ads {
		lf := &pb.LogFilter{AppDef: ad}
		glr.LogFilter = append(glr.LogFilter, lf)
	}
	return
}

// given a comma-delimeted string, will find logappdefs
func filterToApps() []*pb.LogAppDef {
	ctx := authremote.Context()
	x, xe := logServer.GetApps(ctx, &common.Void{})
	utils.Bail("Failed to get apps", xe)
	var res []*pb.LogAppDef
	for _, ld := range x.AppDef {
		if isAppMatch(ld) {
			fmt.Printf("Repository: \"%s\", Application: \"%s\", Namespace: \"%s\",  Groupname: \"%s\"\n", ld.Repository, ld.Appname, ld.Namespace, ld.Groupname)
			res = append(res, ld)
		}
	}
	return res
}

func isAppMatch(def *pb.LogAppDef) bool {
	ms := strings.Split(match, ",")
	for _, s := range ms {
		if strings.Contains(def.Repository, s) {
			return true
		}
	}
	return false
}
