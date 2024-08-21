package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hmcalister/TeaTimer/internal/timerdata"
)

var (
	appStyle = lipgloss.NewStyle().Margin(1, 2)
)

