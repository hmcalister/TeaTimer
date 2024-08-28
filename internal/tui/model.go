package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	linkedlist "github.com/hmcalister/Go-DSA/list/LinkedList"
	"github.com/hmcalister/TeaTimer/internal/timerdata"
)

type tickMsg time.Time

type MainModel struct {
	keybindings         *keybindList
	timerManager        *timerdata.TimerManager
	addTimerPopupActive bool
}

func NewMainModel() MainModel {
	keybinds := newKeybindList()

	timerManager := timerdata.NewManager()
	timerManager.NewTimer("A", 60)
	return MainModel{
		keybindings:  keybinds,
		timerManager: timerManager,
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(500*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m MainModel) Init() tea.Cmd {
	return tickCmd()
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var batchedCmds []tea.Cmd

	switch msg := msg.(type) {
	case tickMsg:
		return m, tickCmd()
	case tea.WindowSizeMsg:
		h, v := mainContentStyle.GetFrameSize()
		mainContentStyle = mainContentStyle.Width(max(msg.Width-h, contentMinWidth)).Height(max(msg.Height-v, contentMinHeight))
		progressBarStyle = progressBarStyle.Width(mainContentStyle.GetWidth() - h)
		h, v = popupContentStyle.GetFrameSize()
		popupContentStyle = popupContentStyle.Width(max(msg.Width/2-h, contentMinWidth)).Height(max(msg.Height/2-2*v, contentMinHeight))
		formLabelStyle = formLabelStyle.Width(popupContentStyle.GetWidth()/3 - h)
		formInputStyle = formInputStyle.Width(2*popupContentStyle.GetWidth()/3 - h)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keybindings.exit):
			return m, tea.Quit

		case key.Matches(msg, m.keybindings.quit):
			if m.addTimerPopupActive {
				m.addTimerPopupActive = false
				return m, tickCmd()
			} else {
				return m, tea.Quit
			}

		case key.Matches(msg, m.keybindings.add):
			if !m.addTimerPopupActive {
				m.addTimerPopupActive = true
			}
			return m, nil
		}
	}

	return m, tea.Batch(batchedCmds...)
}

func (m MainModel) View() string {
	m.timerManager.AllTimersMutex.RLock()
	defer m.timerManager.AllTimersMutex.RUnlock()

	renderString := "TIMER APP\n\n"

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
	renderString += lipgloss.JoinVertical(lipgloss.Left, timerVisuals...)

	if m.addTimerPopupActive {
		form := lipgloss.JoinVertical(
			lipgloss.Left,
			"Add New Timer",
			fmt.Sprintf("%d %d", mainContentStyle.GetWidth(), mainContentStyle.GetHeight()),
			fmt.Sprintf("%d %d", popupContentStyle.GetWidth(), popupContentStyle.GetHeight()),
			lipgloss.JoinHorizontal(
				lipgloss.Center,
				formLabelStyle.Render("Timer Name: "),
				formInputStyle.Render("..."),
			),
			lipgloss.JoinHorizontal(
				lipgloss.Center,
				formLabelStyle.Render("Timer Duration (s): "),
				formInputStyle.Render("..."),
			),
		)
		renderString += lipgloss.Place(
			mainContentStyle.GetWidth(),
			mainContentStyle.GetHeight(),
			lipgloss.Center,
			lipgloss.Center,
			popupContentStyle.Render(form),
		)
	}
	return mainContentStyle.Render(renderString)
}
