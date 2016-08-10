// Author: Sean Hamilton <skhamilt@eng.ucsd.edu>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


package main

import (
	"os/user"
	"fmt"
	"os"
	"flag"
	"syscall"
	"time"
	"./types"
)


var running = true
var params []types.TaskParameters

func timertest(tid int, priority int32) {
	// Setup Code
	//TODO: Error checking the range of priority
	// Ensure the underlying OS thread is set SCHED_FIFO and set priority
	var p syscall.SchedParam
	p.X__sched_priority = priority
	//TODO: check error return types
	syscall.SchedSetscheduler(0,syscall.SCHED_FIFO,&p)

	// Test Code
	now := time.Now()
	next := now + params[tid].Interval
}

func main() {
	// Check to ensure this is being run as root
	u,_ := user.Current()
	if u.Uid != "0" {
		fmt.Fprintln(os.Stderr, "cyclictest must be run as root")
		os.Exit(-syscall.EPERM)
	}

	//-l LOOPS --loops=LOOPS     number of loops: default=0(endless)
	var numLoops int
	flag.IntVar(&numLoops, "loops", 0, "number of loops: default=0(endless)")
	//-p PRIO  --priority=PRIO   priority of highest prio thread
	var highestPrio int
	flag.IntVar(&highestPrio, "priority", 0, "priority of highest prio thread")
	//-t [NUM] --threads=NUM     number of threads:
	var numTasks int
	flag.IntVar(&numTasks, "tasks", 1, "number of tasks")
	var distance int
	flag.IntVar(&distance,"distance", 500, "distance of thread intervals in us default=500")

	flag.Parse()
	//TODO: check for valid ranges of flags

	nextPrio := highestPrio
	for i := 0; i < numTasks; i++ {
		go timertest(i, nextPrio)
		nextPrio += distance
	}

	for running {
		for i := 0; i < numTasks; i++ {
			//stats[i].print()
			time.Sleep(10*time.Microsecond)
		}
	}

}
