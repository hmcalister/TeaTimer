package tui

import (
	"slices"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
	linkedlist "github.com/hmcalister/Go-DSA/list/LinkedList"
	"github.com/hmcalister/TeaTimer/internal/timerdata"
)

func (m AppModel) renderMainPage() string {
	m.timerManager.AllTimersMutex.RLock()
	defer m.timerManager.AllTimersMutex.RUnlock()

	titleContent := titleContentStyle.Render("TIMER APP")

	progressBar := progress.New(progress.WithDefaultGradient())
	progressBar.ShowPercentage = false
	progressBar.Width = progressBarStyle.GetWidth()
	timerVisuals := make([]string, 0)
	linkedlist.ForwardApply(m.timerManager.AllTimers, func(timer *timerdata.TimerData) {
		timerVisuals = append(timerVisuals, lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.JoinHorizontal(
				lipgloss.Center,
				timer.Name,
				": ",
				timer.GetStatusAsString(),
			),
			lipgloss.JoinHorizontal(
				lipgloss.Center,
				progressBar.ViewAs(timer.GetProgressProportion()),
			),
			"\n",
		))
	})
	mainContent := mainContentStyle.Render(lipgloss.JoinVertical(lipgloss.Left, timerVisuals...))

	return lipgloss.JoinVertical(lipgloss.Center, titleContent, mainContent)
}

func (m AppModel) renderAddTimerPage() string {

	titleContent := titleContentStyle.Render("ADD TIMER")

	formInputFocused := formInputStyle.Foreground(lipgloss.Color("205"))
	formInputBlurred := formInputStyle.Foreground(lipgloss.Color("240"))
	formInputStyles := slices.Repeat([]lipgloss.Style{formInputBlurred}, len(m.addTimerInputs))
	formInputStyles[m.focusIndex] = formInputFocused
	mainContent := lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinVertical(
			lipgloss.Center,
			formLabelStyle.Render("Timer Name: "),
			formInputStyles[0].Render(m.addTimerInputs[0].View()),
		),
		"\n",
		lipgloss.JoinVertical(
			lipgloss.Center,
			formLabelStyle.Render("Timer Duration (s): "),
			formInputStyles[1].Render(m.addTimerInputs[1].View()),
		),
		"\n",
	)
	mainContent = mainContentStyle.Render(mainContent)

	return lipgloss.JoinVertical(lipgloss.Center, titleContent, mainContent)
}
