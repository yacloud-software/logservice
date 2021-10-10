package main

import (
	"errors"
	"flag"
	"fmt"
	"golang.conradwood.net/go-easyops/logger"
	"io"
	"os"
	"os/exec"
	"syscall"
)

// static variables for flag parser
var (
	log_status = flag.String("status", "", "The status string to log")
	app_name   = flag.String("appname", "", "The name of the application to log")
	repo       = flag.Uint64("repository", 0, "The name of the repository to log")
	groupname  = flag.String("groupname", "", "The name of the group to log")
	namespace  = flag.String("namespace", "", "the namespace to log")
	deplid     = flag.String("deploymentid", "", "The deployment id to log")
	sid        = flag.String("startupid", "", "The startup id to log")
	is_running = false
)

type com struct {
	Cmd    *exec.Cmd
	Name   string
	Paras  []string
	Stdout io.ReadCloser
	Stderr io.ReadCloser
	code   int
	logger *logger.AsyncLogQueue
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
	name := paras[0]
	if len(paras) > 1 {
		paras = paras[1:]
	}
	cmd := com{Name: name, Paras: paras}
	err := run(&cmd)
	if cmd.logger != nil {
		s := fmt.Sprintf("normal shutdown")
		if err != nil {
			s = fmt.Sprintf("shutdown with %s", err)
		}
		cmd.logger.LogCommandStdout(s, "TERMINATED")
		e := cmd.logger.Close(0)
		if e != nil {
			fmt.Printf("Failed to flush logs: %s\n", e)
		}
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
	lineOut := new(LineReader)
	buf := make([]byte, 2)
	for {
		ct, err := cmd.Stdout.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Failed to read command output: %s\n", err)
			}
			break
		}
		line := lineOut.gotBytes(buf, ct)
		if line != "" {
			cmd.checkLogger()
			fmt.Printf("STDOUT: \"%s\"\n", line)
			cmd.logger.LogCommandStdout(line, "EXECUSER")
		}
	}
	err := cmd.Cmd.Wait()
	if cmd.logger == nil {
		fmt.Printf("No logger to close\n")
	} else {
		cmd.logger.Close(0)
	}
	return err
}

// wait for stderr
func waitForStderr(cmd *com) {
	lineOut := new(LineReader)
	buf := make([]byte, 2)
	for {
		ct, err := cmd.Stderr.Read(buf)
		if err != nil {
			if is_running {
				continue
			}
			break
		}
		line := lineOut.gotBytes(buf, ct)
		if line != "" {
			cmd.checkLogger()
			fmt.Printf("STDERR: \"%s\"\n", line)
			cmd.logger.LogCommandStdout(line, "EXECUSER")
		}
	}
}

func (c *com) checkLogger() {
	if c.logger != nil {
		return
	}
	l, err := logger.NewAsyncLogQueue(*app_name, *repo, *groupname, *namespace, *deplid)
	if err != nil {
		fmt.Printf("Failed to initialize logger! %s\n", err)
		// it's the logger-wrapper! what is it supposed to do if not log-wrap?
		os.Exit(10)
	} else {
		c.logger = l
	}
}
