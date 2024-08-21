package timerdata

type TimerData struct {
	Name              string
	TimerState        TimerStateEnum
	InitialDuration   int
	RemainingDuration int
	UpdateChannel     chan TimerUpdateMessageEnum
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
