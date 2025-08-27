// client create: LogServiceClient
/*
  Created by /home/cnw/devel/go/yatools/src/golang.yacloud.eu/yatools/protoc-gen-cnw/protoc-gen-cnw.go
*/

/* geninfo:
   filename  : golang.conradwood.net/apis/logservice/logservice.proto
   gopackage : golang.conradwood.net/apis/logservice
   importname: ai_0
   clientfunc: GetLogService
   serverfunc: NewLogService
   lookupfunc: LogServiceLookupID
   varname   : client_LogServiceClient_0
   clientname: LogServiceClient
   servername: LogServiceServer
   gsvcname  : logservice.LogService
   lockname  : lock_LogServiceClient_0
   activename: active_LogServiceClient_0
*/

package logservice

import (
   "sync"
   "golang.conradwood.net/go-easyops/client"
)
var (
  lock_LogServiceClient_0 sync.Mutex
  client_LogServiceClient_0 LogServiceClient
)

func GetLogClient() LogServiceClient { 
    if client_LogServiceClient_0 != nil {
        return client_LogServiceClient_0
    }

    lock_LogServiceClient_0.Lock() 
    if client_LogServiceClient_0 != nil {
       lock_LogServiceClient_0.Unlock()
       return client_LogServiceClient_0
    }

    client_LogServiceClient_0 = NewLogServiceClient(client.Connect(LogServiceLookupID()))
    lock_LogServiceClient_0.Unlock()
    return client_LogServiceClient_0
}

func GetLogServiceClient() LogServiceClient { 
    if client_LogServiceClient_0 != nil {
        return client_LogServiceClient_0
    }

    lock_LogServiceClient_0.Lock() 
    if client_LogServiceClient_0 != nil {
       lock_LogServiceClient_0.Unlock()
       return client_LogServiceClient_0
    }

    client_LogServiceClient_0 = NewLogServiceClient(client.Connect(LogServiceLookupID()))
    lock_LogServiceClient_0.Unlock()
    return client_LogServiceClient_0
}

func LogServiceLookupID() string { return "logservice.LogService" } // returns the ID suitable for lookup in the registry. treat as opaque, subject to change.

func init() {
   client.RegisterDependency("logservice.LogService")
   AddService("logservice.LogService")
}
