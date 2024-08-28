package tui

import "github.com/charmbracelet/lipgloss"

const (
	contentMinWidth  = 40
	contentMinHeight = 5
)

var (
	mainContentStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				Width(contentMinWidth).
				Height(contentMinHeight).
				Padding(0, 5)

	progressBarStyle = lipgloss.NewStyle().
				Width(contentMinWidth)

	popupContentStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				Align(lipgloss.Left).
				Width(contentMinWidth/2).
				Height(contentMinHeight/2).
				Padding(0, 5)

	formLabelStyle = lipgloss.NewStyle().
			Width(contentMinWidth / 3).
			Align(lipgloss.Left).
			MarginRight(5)

	formInputStyle = lipgloss.NewStyle().
			Width(contentMinWidth/6).
			Border(lipgloss.RoundedBorder()).
			Align(lipgloss.Left).
			Padding(0, 5)
)
