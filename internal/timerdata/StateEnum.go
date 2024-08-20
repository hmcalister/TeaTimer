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

