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

