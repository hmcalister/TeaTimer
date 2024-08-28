package timerdata

import (
	"fmt"
	"time"
)

type TimerData struct {
	Name              string
	TimerState        TimerStateEnum
	InitialDuration   int
	RemainingDuration int
	UpdateChannel     chan TimerUpdateMessageEnum
}

func (t *TimerData) GetProgressProportion() float64 {
	return 1 - float64(t.RemainingDuration)/float64(t.InitialDuration)
}

func (t *TimerData) GetRemainingDurationAsString() string {
	remainingDuration := time.Duration(t.RemainingDuration) * time.Second
	days := int(remainingDuration.Hours()) / 24
	hours := int(remainingDuration.Hours()) % 24
	minutes := int(remainingDuration.Minutes()) % 60
	seconds := int(remainingDuration.Seconds()) % 60

	if days > 0 {
		return fmt.Sprintf("%d Days %d Hours %d Minutes %d Seconds", days, hours, minutes, seconds)
	}
	if hours > 0 {
		return fmt.Sprintf("%d Hours %d Minutes %d Seconds", hours, minutes, seconds)
	}
	if minutes > 0 {
		return fmt.Sprintf("%d Minutes %d Seconds", minutes, seconds)
	}
	if seconds > 0 {
		return fmt.Sprintf("%d Seconds", seconds)
	}
	return ""
}

// Function to run concurrently, as `go t.stateMachine`.
//
// Handles all incoming messages, as well as ticking the timer duration as required.
// Concurrent reads shouldn't (?) be an issue here, as being out of sync by one second
// likely won't be an issue... but could introduce locks later if need be.
//
// Takes a channel that is closed on starting, to signal that this timer has started processing states.
func (t *TimerData) stateMachine(startSignal chan interface{}) {
	close(startSignal)
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
			if t.TimerState == TimerStateRunning {
				t.RemainingDuration -= 1
				if t.RemainingDuration <= 0 {
					t.TimerState = TimerStateRinging
					t.RemainingDuration = 0
				}
			}
		}
	} // End Infinite Loop

	t.TimerState = TimerStateFinished
}
