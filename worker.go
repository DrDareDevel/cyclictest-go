package main

import (
	"syscall"
	"runtime"
	"time"
	"fmt"
	"github.com/RedShamilton/cyclictest-go/types"
)

func worker(param *types.TaskParameters, itrs uint) {
	// Setup Code
	// Ensure the underlying OS thread is set SCHED_FIFO and set priority
	var p syscall.SchedParam
	p.X__sched_priority = param.Priority
	//TODO: implement error checking on syscalls
	syscall.SchedSetscheduler(0,syscall.SCHED_FIFO,&p)
        runtime.LockOSThread()
	stats := param.Stats
	stats.Reset()

	// Test Code
	for i := uint(0); i < itrs && running; i++ {
                next := time.Now().Add(param.Interval)
                latency := time.WaitUntil(next)
		stats.Update(latency)
        	
		if histogram {
                	fmt.Fprintf(histfile, "%v\t%v\n",i,latency.Nanoseconds())
        	}
	}

        running = false
        wg.Done()
}

