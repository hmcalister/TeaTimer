package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hmcalister/TeaTimer/internal/timerdata"
)

type tickMsg time.Time

type AppModel struct {
	keybindings  *keybindList
	timerManager *timerdata.TimerManager
	viewState    viewStateEnum
}

func NewMainModel() AppModel {
	keybinds := newKeybindList()

	timerManager := timerdata.NewManager()
	timerManager.NewTimer("A", 60)
	timerManager.NewTimer("My Timer with a Cool Name", 888888)
	timerManager.NewTimer("My less cool timer", 222)
	timerManager.NewTimer("Genuinely a disappointment", 123)
	timerManager.NewTimer("OFF DA PAGE", 15)
	return AppModel{
		keybindings:  keybinds,
		timerManager: timerManager,
		viewState:    mainTimerPage,
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(500*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m AppModel) Init() tea.Cmd {
	return tickCmd()
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var batchedCmds []tea.Cmd

	switch msg := msg.(type) {
	case tickMsg:
		return m, tickCmd()
	case tea.WindowSizeMsg:
		h, _ := titleContentStyle.GetFrameSize()
		titleContentStyle = titleContentStyle.Width(max(msg.Width-h, contentMinWidth))
		h, v := mainContentStyle.GetFrameSize()
		mainContentStyle = mainContentStyle.Width(max(msg.Width-h, contentMinWidth)).Height(max(msg.Height-v-titleContentStyle.GetHeight()-2, contentMinHeight))
		progressBarStyle = progressBarStyle.Width(mainContentStyle.GetWidth() - h)
		formLabelStyle = formLabelStyle.Width(max(msg.Width/2-h, contentMinWidth/3))
		formInputStyle = formInputStyle.Width(max(msg.Width/2-h, contentMinWidth/3))

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keybindings.exit):
			return m, tea.Quit

		case key.Matches(msg, m.keybindings.quit):
			if m.viewState == mainTimerPage {
				return m, tea.Quit
			} else {
				m.viewState = mainTimerPage
				return m, tickCmd()
			}

		case key.Matches(msg, m.keybindings.add):
			if m.viewState == mainTimerPage {
				m.viewState = addTimerPage
			}
			return m, nil
		}
	}

	return m, tea.Batch(batchedCmds...)
}

func (m AppModel) View() string {
	switch m.viewState {
	case mainTimerPage:
		return m.renderMainPage()
	case addTimerPage:
		return m.renderAddTimerPage()
	}
	return ""
}
