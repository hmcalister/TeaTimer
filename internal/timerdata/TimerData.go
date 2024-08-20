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

// Function to run concurrently, as `go t.stateMachine`.
//
// Handles all incoming messages, as well as ticking the timer duration as required.
// Concurrent reads shouldn't (?) be an issue here, as being out of sync by one second
// likely won't be an issue... but could introduce locks later if need be.
//
// Once the timer is deleted, this method cleans up any remaining channels.
func (t *TimerData) stateMachine() {
	for updateMessage := range t.UpdateChannel {
		switch updateMessage {
		case UpdateMessagePause:
			if t.TimerState == TimerStateRunning {
				t.TimerState = TimerStatePaused
			}

		case UpdateMessageUnpause:
			if t.TimerState == TimerStatePaused {
				t.TimerState = TimerStateRunning
			}

		case UpdateMessageStop:
			t.RemainingDuration = 0
			t.TimerState = TimerStateFinished

		case UpdateMessageRestart:
			t.RemainingDuration = t.InitialDuration
			t.TimerState = TimerStateRunning

		case UpdateMessageTick:
			fmt.Printf("%v CONSUMED\n", t.Name)
			if t.TimerState == TimerStateRunning {
				t.RemainingDuration -= 1
				fmt.Printf("\tDURATION: %v\n", t.RemainingDuration)
				if t.RemainingDuration <= 0 {
					fmt.Printf("\tRINGING\n")
					t.TimerState = TimerStateRinging
					t.RemainingDuration = 0
				}
			}
		}
	} // End Infinite Loop

	t.TimerState = TimerStateFinished
	fmt.Printf("%v CLOSED\n", t.Name)
}
