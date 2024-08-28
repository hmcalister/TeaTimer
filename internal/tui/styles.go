package tui

import "github.com/charmbracelet/lipgloss"

var (
	mainContent = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Width(80).
			Height(30).
			Padding(0, 5)

	popupContent = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Width(40).
			Height(20).
			Padding(0, 5)

	formLabelStyle = lipgloss.NewStyle().
			Width(40).
			Align(lipgloss.Left).
			MarginRight(5)

	formInputStyle = lipgloss.NewStyle().
		// Width(20).
		Align(lipgloss.Right)
)
