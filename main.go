// Author: Sean Hamilton <skhamilt@eng.ucsd.edu>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


package main

import (
	"os/user"
	"fmt"
	"log"
	"os"
	"flag"
	"sync"
	"syscall"
	"time"
	"github.com/RedShamilton/cyclictest-go/types"
)

var wg sync.WaitGroup
var running = true
var histogram = false
var histfile *os.File
var params []types.TaskParameters

func main() {
	var err error 
	
	// Check to ensure this is being run as root
	u,_ := user.Current()
	if u.Uid != "0" {
		fmt.Fprintln(os.Stderr, "cyclictest must be run as root")
		os.Exit(int(syscall.EPERM))
	}

	var numLoops uint
	flag.UintVar(&numLoops, "l", 1000, "number of `loops`")

	var priority int
	flag.IntVar(&priority, "p", 0, "priority for highest `priority` task")

	var numTasks uint
	flag.UintVar(&numTasks, "t", 1, "number of `tasks`")

	var distance time.Duration
	flag.DurationVar(&distance,"d", 500*time.Microsecond, "`distance` of task intervals")

	var interval time.Duration
	flag.DurationVar(&interval,"i", 1000*time.Microsecond, "base `interval` of task")

	var histfilename string
	flag.StringVar(&histfilename, "H", "", "dump a latency histogram to `file` after the run")

	flag.Parse()
	//TODO: check for valid ranges of flags

	// Should we print out histogram data?
	if histfilename != "" {
		histogram = true
		histfile, err = os.Create(histfilename)
		if err != nil {
			log.Fatal(err)
		}
	}

        wg.Add(int(numTasks))
	nextInterval := interval
	nextPrio := priority
	params = make([]types.TaskParameters, numTasks)
	for i := uint(0); i < numTasks; i++ {
		params[i].Init(i,nextInterval,int32(nextPrio),new(types.TaskStatistics))
		nextPrio -= 1
		nextInterval += distance
		go worker(&params[i], numLoops)
	}

        wg.Wait()

	for i := uint(0); i < numTasks; i++ {
		params[i].PrintResults()
	}
}
