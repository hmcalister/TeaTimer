package timerdata

import "time"

type TimerData struct {
	Name      string
	StartTime time.Time
	Duration  time.Duration
	Repeat    bool
}

