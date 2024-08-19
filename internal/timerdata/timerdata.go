package timerdata

import "time"

// TODO: How to handle:
// - Start
// - Stop
// - Pause (Find the amount of duration elapsed, and keep that?)
// - Delete (e.g. garbage collection / channel closing)
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
