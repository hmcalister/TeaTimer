// Code generated by "stringer -type=TimerUpdateMessageEnum"; DO NOT EDIT.

package timerdata

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UpdateMessagePause-0]
	_ = x[UpdateMessageUnpause-1]
	_ = x[UpdateMessageStop-2]
	_ = x[UpdateMessageRestart-3]
}

const _TimerUpdateMessageEnum_name = "UpdateMessagePauseUpdateMessageUnpauseUpdateMessageStopUpdateMessageRestart"

var _TimerUpdateMessageEnum_index = [...]uint8{0, 18, 38, 55, 75}

func (i TimerUpdateMessageEnum) String() string {
	if i < 0 || i >= TimerUpdateMessageEnum(len(_TimerUpdateMessageEnum_index)-1) {
		return "TimerUpdateMessageEnum(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TimerUpdateMessageEnum_name[_TimerUpdateMessageEnum_index[i]:_TimerUpdateMessageEnum_index[i+1]]
}
