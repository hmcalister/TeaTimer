package timerdata

// Enumerate the states of a TimerData
//
//go:generate stringer -type=TimerStateEnum
type TimerStateEnum int

const (
	// State of a timer that is currently active.
	// A timer in this state is currently counting down.
	TimerStateRunning TimerStateEnum = iota

	// State of a timer that is currently paused.
	// A timer in this state is not currently counting down.
	TimerStatePaused TimerStateEnum = iota

	// State of a timer that is currently ringing.
	// A timer in this state has zero remaining duration,
	// and the timer is ringing, waiting to be dismissed
	TimerStateRinging TimerStateEnum = iota

	// State of a timer that has finished.
	// A timer in this state has zero remaining duration,
	// although the timer is not ringing.
	TimerStateFinished TimerStateEnum = iota
)

// Enumerate the messages that can be sent to a timer.
//
//go:generate stringer -type=TimerUpdateMessageEnum
type TimerUpdateMessageEnum int

const (
	// A message to pause the timer.
	// This message is only valid for a timer in state TimerStateRunning,
	// and sends the timer to TimerStatePaused.
	// All other states will ignore this message.
	UpdateMessagePause TimerUpdateMessageEnum = iota

	// A message to unpause the timer.
	// This message is only valid for a timer in state TimerStatePaused,
	// and sends the timer to TimerStateRunning.
	UpdateMessageUnpause TimerUpdateMessageEnum = iota

	// A message to stop the timer.
	// This message is valid for timers in states TimerStateRunning,
	// TimerStatePaused, or TimerStateRinging.
	// Sends timer to TimerStateFinished.
	UpdateMessageStop TimerUpdateMessageEnum = iota

	// A message to restart the timer.
	// This message is valid for timers in state TimerStateFinished.
	// Sends timers to state TimerStateRunning, with a refreshed remaining duration.
	UpdateMessageRestart TimerUpdateMessageEnum = iota

	// A single tick event, decrements the remaining duration by one.
	UpdateMessageTick TimerUpdateMessageEnum = iota
)
