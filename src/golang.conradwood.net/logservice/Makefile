.PHONY:	client server wrapper
server:
	cd server && go install ${LDFLAGS} logservice-server.go `ls -1 *.go|grep -v logservice-server.go`

client-glog:
	cd glog && go install ${LDFLAGS} glog.go helpers.go
client:
	cd client && go install ${LDFLAGS} logservice-client.go 
wrapper:
	cd wrapper && go install ${LDFLAGS} logger-wrapper.go linereader.go

all: client client-glog server wrapper

