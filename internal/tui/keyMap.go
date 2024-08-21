package tui

import "github.com/charmbracelet/bubbles/key"

type keybindList struct {
	quit    key.Binding
	suspend key.Binding
	add     key.Binding
}

func newKeybindList() *keybindList {
	return &keybindList{
		quit: key.NewBinding(
			key.WithKeys("q"),
			key.WithHelp("q", "Quit"),
		),
		suspend: key.NewBinding(
			key.WithKeys("ctrl+z"),
			key.WithHelp("ctrl+z", "Suspend"),
		),
		add: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "New Timer"),
		),
	}
}
