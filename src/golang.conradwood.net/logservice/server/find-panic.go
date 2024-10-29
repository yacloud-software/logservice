package main

import (
	"fmt"
	pb "golang.conradwood.net/apis/logservice"
	"golang.conradwood.net/go-easyops/authremote"
	pn "golang.yacloud.eu/apis/crashanalyser"
	"time"
)

func checkPanic(ad *pb.LogAppDef, lines []string) {

	// call the panic request thingie
	pdr := pn.AnalyseLogRequest{
		Repository: fmt.Sprintf("%d", ad.RepoID),
		Namespace:  ad.Namespace,
		Groupname:  ad.Groupname,
		Appname:    ad.Appname,
		Build:      0, // where do we get that from?
	}
	for _, line := range lines {
		ll := &pn.LogLine{
			Host:    "logservice:unknown",
			Occured: uint64(time.Now().Unix()),
			Line:    line,
			Status:  "logservice:unknown",
		}
		pdr.Lines = append(pdr.Lines, ll)
	}
	ctx := authremote.Context()
	pn.GetCrashAnalyserClient().AnalyseLogs(ctx, &pdr)
}
