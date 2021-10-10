package main

import (
	"flag"
	"fmt"
	"golang.conradwood.net/go-easyops/prometheus"
	"golang.conradwood.net/go-easyops/utils"
	"os"
	"sync"
)

var (
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

	fmt.Printf("[rotate] rotating...\n")
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

}
