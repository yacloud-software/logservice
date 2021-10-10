.PHONY:	client server wrapper
server:
	cd server && go install logservice-server.go `ls -1 *.go|grep -v logservice-server.go`

client-glog:
	cd glog && go install glog.go helpers.go
client:
	cd client && go install logservice-client.go 
wrapper:
	cd wrapper && go install logger-wrapper.go linereader.go

all: client client-glog server wrapper

