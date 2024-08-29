package tui

import "github.com/charmbracelet/bubbles/key"

type keybindList struct {
	exit       key.Binding
	quit       key.Binding
	add        key.Binding
	scrollUp   key.Binding
	scrollDown key.Binding
	submit     key.Binding
}

func newKeybindList() *keybindList {
	return &keybindList{
		exit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("q", "Exit"),
		),
		quit: key.NewBinding(
			key.WithKeys("q"),
			key.WithHelp("q", "Quit"),
		),
		add: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "New Timer"),
		),
		scrollUp: key.NewBinding(
			key.WithKeys("k", "up", "shift+tab"),
			key.WithHelp("up", "Scroll Up"),
		),
		scrollDown: key.NewBinding(
			key.WithKeys("j", "down", "tab"),
			key.WithHelp("down", "Scroll Down"),
		),
		submit: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Submit"),
		),
	}
}
