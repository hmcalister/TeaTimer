package timerdata

import "time"

type TimerData struct {
	Name      string
	StartTime time.Time
	Duration  time.Duration
	Repeat    bool
}

func NewTimer() *TimerData {
	return nil
}

func (t *TimerData) StartTimer() {

}

func (t *TimerData) StopTimer() {

}
