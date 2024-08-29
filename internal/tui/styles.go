package tui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	contentMinWidth  = 40
	contentMinHeight = 5
)

var (
	titleContentStyle = lipgloss.NewStyle().
				Align(lipgloss.Center).
				Border(lipgloss.RoundedBorder()).
				Width(contentMinWidth).
				Height(1).
				Padding(0, 5)

	mainContentStyle = lipgloss.NewStyle().
				Align(lipgloss.Center).
				Border(lipgloss.RoundedBorder()).
				Width(contentMinWidth).
				Height(contentMinHeight).
				Padding(0, 5)

	progressBarStyle = lipgloss.NewStyle().
				Width(contentMinWidth)

	formLabelStyle = lipgloss.NewStyle().
			Align(lipgloss.Left).
			Width(contentMinWidth / 4).
			MarginRight(5)

	formInputStyle = lipgloss.NewStyle().
			Align(lipgloss.Left).
			Border(lipgloss.RoundedBorder()).
			Width(contentMinWidth/4).
			Padding(0, 5)
)
