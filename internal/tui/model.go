package tui

import (
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
	allKeybinds := func() []key.Binding {
		return []key.Binding{
			keybinds.add,
			keybinds.suspend,
			// keybinds.quit,
		}
	}
	timersList := list.New(make([]list.Item, 0), list.NewDefaultDelegate(), 0, 0)
	timersList.AdditionalShortHelpKeys = allKeybinds
	timersList.AdditionalFullHelpKeys = allKeybinds
	return MainModel{
		keybindings:  keybinds,
		timersList:   timersList,
		timerManager: timerdata.NewManager(),
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
	switch msg := msg.(type) {
	case tickMsg:
		return m, tickCmd()
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.timersList.SetSize(msg.Width-h, msg.Height-v)

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

		case key.Matches(msg, m.keybindings.suspend):
			return m, tea.Suspend

		case key.Matches(msg, m.keybindings.add):
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.timersList, cmd = m.timersList.Update(msg)

	return m, cmd
}

func (m MainModel) View() string {
	return appStyle.Render(m.timersList.View())
}
