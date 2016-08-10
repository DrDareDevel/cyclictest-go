package types

import "time"

type TaskParameters struct {
	Priority int
	//policy int
	//mode int
	//timermode int
	//signal int
	//clock int
	//max_cycles uint64
	Stats []TaskStatistics
	//bufmsk int
	Interval time.Duration
	//cpu int
	//node int
	Id int
	//msr_fd int
}

