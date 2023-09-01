package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/logservice"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"syscall"
)

// static variables for flag parser
var (
	log_status = flag.String("status", "", "The status string to log")
	is_running = false
	ls         logservice.LogServiceClient
)

type com struct {
	Cmd       *exec.Cmd
	Name      string
	Paras     []string
	Stdout    io.ReadCloser
	Stderr    io.ReadCloser
	code      int
	la        *logservice.LogAppDef
	remainder []byte
	lck       sync.Mutex
}

func (c *com) toString() string {
	return c.Name
}

func bail(err error, msg string) {
	if err == nil {
		return
	}
	fmt.Printf("%s: %s\n", msg, err)
	os.Exit(10)
}

func main() {
	flag.Parse()
	// redirect stderr to stdout (to capture panics)
	if 1 == 2 {
		syscall.Dup2(int(os.Stdout.Fd()), int(os.Stderr.Fd()))
	}
	paras := flag.Args()
	if len(paras) < 1 {
		fmt.Printf("Usage: [name] [parameter1..n]\n")
		fmt.Printf("Error: Missing name and/or parameters\n")
		os.Exit(10)
	}
	ls = logservice.GetLogServiceClient()
	name := paras[0]
	if len(paras) > 1 {
		paras = paras[1:]
	}
	short_name := filepath.Base(name)
	cmd := &com{
		Name:  name,
		Paras: paras,
		la: &logservice.LogAppDef{
			Appname:      short_name,
			Repository:   short_name,
			Groupname:    short_name,
			Namespace:    short_name,
			DeploymentID: utils.RandomString(32),
			StartupID:    utils.RandomString(32),
			RepoID:       0,
			BuildID:      0,
		},
	}
	err := run(cmd)
	ctx := authremote.Context()
	cmd.Write([]byte("TERMINATED\n"))
	_, e := ls.CloseLog(ctx, &logservice.CloseLogRequest{AppDef: cmd.la, ExitCode: 0})
	if e != nil {
		fmt.Printf("Failed to flush logs: %s\n", e)
	}

	if err != nil {
		fmt.Printf("Command %s returned with error: %s (thus logger-wrapper exiting with code 10)\n", cmd.Name, err)
		os.Exit(10)
	}
	os.Exit(0)
}
func run(cmd *com) error {
	var err error
	c := exec.Command(cmd.Name, cmd.Paras...)
	cmd.Cmd = c
	cmd.Stdout, err = cmd.Cmd.StdoutPipe()
	if err != nil {
		s := fmt.Sprintf("Could not get cmd stdout: %s\n", err)
		return errors.New(s)
	}
	cmd.Stderr, err = cmd.Cmd.StderrPipe()
	if err != nil {
		s := fmt.Sprintf("Could not get cmd stderr: %s\n", err)
		return errors.New(s)
	}
	fmt.Printf("Starting Command: %s\n", cmd.Name)
	is_running = true
	err = cmd.Cmd.Start()
	if err != nil {
		fmt.Printf("Command: %v failed\n", cmd)
		return err
	}
	go waitForStderr(cmd)
	// reap children...
	err = waitForCommand(cmd)
	is_running = false
	return err
}

// async, whenever a process exits...
func waitForCommand(cmd *com) error {
	buf := make([]byte, 1024)
	for {
		ct, err := cmd.Stdout.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Failed to read command output: %s\n", err)
			}
			break
		}
		b := buf[:ct]
		fmt.Print(string(b))
		cmd.Write(b)
	}
	err := cmd.Cmd.Wait()
	return err
}

// wait for stderr
func waitForStderr(cmd *com) {
	buf := make([]byte, 1024)
	for {
		ct, err := cmd.Stderr.Read(buf)
		if err != nil {
			if is_running {
				continue
			}
			break
		}
		b := buf[:ct]
		fmt.Print(string(b))
		cmd.Write(b)
	}
}
func (c *com) Write(buf []byte) {
	c.lck.Lock()
	bytes_to_sent := append(c.remainder, buf...)
	idx := bytes.LastIndexByte(bytes_to_sent, '\n')
	if idx == -1 {
		// no \n in entire buffer, queue it a bit, do nothing
		c.remainder = bytes_to_sent
		c.lck.Unlock()
		return
	}
	if idx == len(bytes_to_sent) {
		// cr at the end of buf
		c.remainder = c.remainder[:0]
	} else {
		c.remainder = bytes_to_sent[idx+1:]
		bytes_to_sent = bytes_to_sent[:idx+1] // include the '\n'
	}

	c.lck.Unlock()
	lr := &logservice.LogRequest{
		AppDef: c.la,
		Lines: []*logservice.LogLine{
			&logservice.LogLine{Message: bytes_to_sent},
		},
	}
	ctx := authremote.Context()
	_, err := ls.LogCommandStdout(ctx, lr)
	if err != nil {
		fmt.Printf("FAILED TO LOG: %s\n", utils.ErrorString(err))
	}
}
