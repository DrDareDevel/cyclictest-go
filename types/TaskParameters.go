package types

import "fmt"
import "io"
import "os"
import "time"

type TaskParameters struct {
	Id uint
	Interval time.Duration
	Priority int32
	Stats *TaskStatistics
}

func (tp *TaskParameters) Init(id uint, ival time.Duration, pri int32, stats *TaskStatistics) {
	tp.Id = id
	tp.Interval = ival
	tp.Priority = pri
	tp.Stats = stats
}

func (tp *TaskParameters) PrintResults() {
  tp.FprintResults(os.Stdout)
}

//T: 0 (13525) P: 0 I:1000 C:   3911 Min:      1 Act:    1 Avg:    1 Max:       2
func (tp *TaskParameters) FprintResults(w io.Writer) {
  fmt.Fprintf(w, "T:%3v P:%2v I:%11v ", tp.Id, tp.Priority, tp.Interval)
  tp.Stats.FprintResults(w)
}
