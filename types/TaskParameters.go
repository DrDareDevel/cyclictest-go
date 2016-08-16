package types

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
