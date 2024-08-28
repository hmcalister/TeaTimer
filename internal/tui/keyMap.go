package tui

import "github.com/charmbracelet/bubbles/key"

type keybindList struct {
	exit key.Binding
	quit key.Binding
	add  key.Binding
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
	}
}
