package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/linux"
	"golang.conradwood.net/go-easyops/prometheus"
	"golang.conradwood.net/go-easyops/utils"
)

const (
	MAX_LOG_FILES = 9
)

var (
	bzip_chan = make(chan bool, 100)
	bzip_lock sync.Mutex
	sizeGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "logservice_filesize",
			Help: "V=1 UNIT=decbytes DESC=size of logfiles",
		},
		[]string{"file"},
	)
	rotate_lock sync.Mutex
	maxsize     = flag.Int("max_size_mb", 1024, "maximum byte size (in Megabytes) of the logfile before being rotated")
)

func init() {
	prometheus.MustRegister(sizeGauge)
}
func rotate_loop() {
	for {
		utils.RandomStall(2)
	}

}

func rotate() {
	fi, err := os.Stat(*logfileName)
	if err != nil {
		fmt.Printf("[rotate] Unable to stat %s: %s\n", *logfileName, err)
		return
	}
	// get the size
	size := fi.Size()
	sizeGauge.With(prometheus.Labels{"file": "main"}).Set(float64(size))
	maxmb := (int64(*maxsize) * 1024 * 1024)
	if size < maxmb {
		return
	}
	fmt.Printf("[rotate] Filesize: %d (max: %d), wait for lock\n", size, maxmb)
	rotate_lock.Lock()
	defer rotate_lock.Unlock()
	fmt.Printf("[rotate] Filesize: %d (max: %d), got lock\n", size, maxmb)

	// check again once we got the lock
	fi, err = os.Stat(*logfileName)
	if err != nil {
		fmt.Printf("[rotate] Unable to stat %s: %s\n", *logfileName, err)
		return
	}
	// get the size
	size = fi.Size()
	maxmb = (int64(*maxsize) * 1024 * 1024)
	if size < maxmb {
		return
	}
	bzip_lock.Lock()
	defer bzip_lock.Unlock()

	fmt.Printf("[rotate] rotating...\n")
	for i := (MAX_LOG_FILES - 1); i > 0; i-- {
		shift_filename := fmt.Sprintf("%s.%d", *logfileName, i)
		if utils.FileExists(shift_filename) {
			newfilename := fmt.Sprintf("%s.%d", *logfileName, (i + 1))
			err = os.Rename(shift_filename, newfilename)
			if err != nil {
				fmt.Printf("failed to shift: %s\n", err)
			}
		}

		// shift the .bz2 files, if any
		shift_filename = fmt.Sprintf("%s.%d.bz2", *logfileName, i)
		if utils.FileExists(shift_filename) {
			newfilename := fmt.Sprintf("%s.%d.bz2", *logfileName, (i + 1))
			err = os.Rename(shift_filename, newfilename)
			if err != nil {
				fmt.Printf("failed to shift: %s\n", err)
			}
		}

	}
	newFilename := fmt.Sprintf("%s.1", *logfileName)
	os.Remove(newFilename)
	err = os.Rename(*logfileName, newFilename)
	if err != nil {
		fmt.Printf("[rotate] failed to rename %s to %s: %s\n", *logfileName, newFilename, err)
		return
	}
	if logfile != nil {
		logfile.Close()
		logfile = nil
	}
	logfile, err = utils.OpenWriteFile(*logfileName)
	if err != nil {
		fmt.Printf("[rotate] Failed to open logfile: %s\n", err)
		return
	}
	bzip_chan <- true

}

func bzipp_loop() {
	for {
		<-bzip_chan
		bzipper()
	}
}
func bzipper() {
	bzip_lock.Lock()
	defer bzip_lock.Unlock()
	for i := (MAX_LOG_FILES - 1); i > 0; i-- {
		to_zip_file := fmt.Sprintf("%s.%d", *logfileName, i)
		if !utils.FileExists(to_zip_file) {
			continue
		}
		zip_file := fmt.Sprintf("%s.%d.bz2", *logfileName, i)
		if utils.FileExists(zip_file) {
			ctx := authremote.Context()
			utils.LogFault(ctx, "logfiles out of order", fmt.Sprintf("Zipping file %s: %s exists already", zip_file, to_zip_file))
			break
		}
		l := linux.New()
		l.SetMaxRuntime(time.Duration(10) * time.Minute)
		cmd := []string{"/usr/bin/bzip2", to_zip_file}
		out, err := l.SafelyExecute(cmd, nil)
		if err != nil {
			fmt.Printf("Failed to pbzip: %s\n%s\n", err, string(out))
			break
		}
	}
}
