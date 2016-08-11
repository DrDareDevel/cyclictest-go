package types

import "time"

type TaskParameters struct {
	Id uint
	Interval time.Duration
	Priority uint
	Stats *TaskStatistics
}

func (tp *TaskParameters) Init(id int, ival time.Duration, pri uint, stats *TaskStatistics) {
	tp.Id = id
	tp.Interval = ival
	tp.Priority = pri
	tp.Stats = stats
}
