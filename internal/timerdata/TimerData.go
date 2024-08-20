package timerdata

import (
	"fmt"
)

type TimerData struct {
	Name              string
	TimerState        TimerStateEnum
	InitialDuration   int
	RemainingDuration int
	UpdateChannel     chan TimerUpdateMessageEnum
}

// Creates and returns a new Timer.
//
// When finished with the timer, close the update channel to signal cleanup.
func NewTimer(name string, duration int) *TimerData {
	updateChannel := make(chan TimerUpdateMessageEnum)
	t := &TimerData{
		Name:              name,
		TimerState:        TimerStateRunning,
		InitialDuration:   duration,
		RemainingDuration: duration,
		UpdateChannel:     updateChannel,
	}


	return t
}

