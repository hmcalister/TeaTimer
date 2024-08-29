package tui

//go:generate stringer -type=viewStateEnum
type viewStateEnum int

const (
	mainTimerPage viewStateEnum = iota
	addTimerPage  viewStateEnum = iota
)
