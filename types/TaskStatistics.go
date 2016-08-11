package types

import (
	"fmt"
	"io"
	"time"
	"os"
	"math"
)

type TaskStatistics struct {
	Act        time.Duration
	Count      uint
	Max        time.Duration
	Min        time.Duration

	cumulative time.Duration
}

func (ts *TaskStatistics) Avg() time.Duration {
	return ts.cumulative / ts.Count
}

func (ts *TaskStatistics) Update(val time.Duration) {
	ts.Min    =  min(ts.Min, val)
	ts.Max    =  max(ts.Max, val)
	ts.cumulative += val
	ts.Act    =  val
	ts.Count  += 1
}

func (ts *TaskStatistics) Reset() {
	ts.Min    = 0*time.Second
	ts.Max    = time.Duration(math.MaxInt64)
	ts.cumulative = 0*time.Second
	ts.Act    = 0*time.Second
	ts.Count  = 0
}

func (ts *TaskStatistics) FprintResults(w io.Writer) {
	fmt.Fprintln(w, "Hello")
}

func (ts *TaskStatistics) PrintResults() {
	ts.FprintResults(os.Stdout)
}

func min(d1 time.Duration, d2 time.Duration) time.Duration {
	if d2 < d1 {
		return d2
	}
	return d1
}

func max(d1 time.Duration, d2 time.Duration) time.Duration {
	if d2 > d1 {
		return d2
	}
	return d1
}
