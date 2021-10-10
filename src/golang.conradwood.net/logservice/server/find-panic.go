package main

import (
	"fmt"
	pn "golang.conradwood.net/apis/codeanalyser"
	pb "golang.conradwood.net/apis/logservice"
	"golang.conradwood.net/go-easyops/tokens"
	"time"
)

var (
	ca pn.CodeAnalyserServiceClient
)

func checkPanic(ad *pb.LogAppDef, lines []string) {
	if ca == nil {
		ca = pn.GetCodeAnalyserServiceClient()
	}

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
	ctx := tokens.ContextWithToken()
	ca.AnalyseLogs(ctx, &pdr)
}
