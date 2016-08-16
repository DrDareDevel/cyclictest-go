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
    if ts.Count != 0 {
	return time.Duration(ts.cumulative / time.Duration(ts.Count))
      }

      return 0
}

func (ts *TaskStatistics) Update(val time.Duration) {
	ts.Min    =  min(ts.Min, val)
	ts.Max    =  max(ts.Max, val)
	ts.cumulative += val
	ts.Act    =  val
	ts.Count  += 1
}

func (ts *TaskStatistics) Reset() {
	ts.Max    = 0*time.Second
	ts.Min    = time.Duration(math.MaxInt64)
	ts.cumulative = 0*time.Second
	ts.Act    = 0*time.Second
	ts.Count  = 0
}

func (ts *TaskStatistics) FprintResults(w io.Writer) {
  fmt.Fprintf(w, "C: %8v Min: %11v Act: %11v Avg: %11v Max: %11v\n", ts.Count, ts.Min, ts.Act, ts.Avg(), ts.Max)
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
