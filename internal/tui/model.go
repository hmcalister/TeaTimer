package tui

import (
	"errors"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hmcalister/TeaTimer/internal/timerdata"
)

type tickMsg time.Time

type AppModel struct {
	keybindings  *keybindList
	timerManager *timerdata.TimerManager
	viewState    viewStateEnum

	focusIndex int

	// inputs corresponding to addTimer form
	// First input is timer name, second input is timer duration (expecting integer in seconds)
	addTimerInputs []textinput.Model
}

func NewMainModel() AppModel {
	keybinds := newKeybindList()

	timerManager := timerdata.NewManager()
	timerManager.NewTimer("A", 60)
	timerManager.NewTimer("My Timer with a Cool Name", 888888)
	// timerManager.NewTimer("My less cool timer", 222)
	// timerManager.NewTimer("Genuinely a disappointment", 123)
	// timerManager.NewTimer("OFF DA PAGE", 15)

	addTimerNameInput := textinput.New()
	addTimerNameInput.Placeholder = "Name..."
	addTimerNameInput.CharLimit = 32
	addTimerNameInput.Validate = func(s string) error {
		if len(s) < 1 {
			return errors.New("timer must have name")
		}
		return nil
	}

	addTimerDurationInput := textinput.New()
	addTimerDurationInput.Placeholder = "Duration..."
	addTimerDurationInput.Validate = func(s string) error {
		parsedInt, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		if parsedInt <= 0 {
			return errors.New("timer must have positive duration")
		}
		return nil
	}
	return AppModel{
		keybindings:  keybinds,
		timerManager: timerManager,
		viewState:    mainTimerPage,
		addTimerInputs: []textinput.Model{
			addTimerNameInput,
			addTimerDurationInput,
		},
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(500*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m AppModel) Init() tea.Cmd {
	return tea.Batch(tickCmd(), textinput.Blink)
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
		}
	}

	switch m.viewState {
	case mainTimerPage:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, m.keybindings.quit):
				return m, tea.Quit
			case key.Matches(msg, m.keybindings.add):
				m.viewState = addTimerPage
				m.focusIndex = 0
				m.addTimerInputs[m.focusIndex].Focus()
				return m, nil
			}
		}

	case addTimerPage:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, m.keybindings.quit):
				m.viewState = mainTimerPage
				m.addTimerInputs[m.focusIndex].Blur()
				m.focusIndex = 0
				return m, nil
			case key.Matches(msg, m.keybindings.scrollUp):
				m.addTimerInputs[m.focusIndex].Blur()
				m.focusIndex = max(m.focusIndex-1, 0)
				m.addTimerInputs[m.focusIndex].Focus()
			case key.Matches(msg, m.keybindings.scrollDown):
				m.addTimerInputs[m.focusIndex].Blur()
				m.focusIndex = min(m.focusIndex+1, len(m.addTimerInputs)-1)
				m.addTimerInputs[m.focusIndex].Focus()
			case key.Matches(msg, m.keybindings.submit):
				for _, input := range m.addTimerInputs {
					inputValue := input.Value()
					if err := input.Validate(inputValue); err != nil {
						input.Reset()
					}
				}
				timerName := m.addTimerInputs[0].Value()
				timerDuration := m.addTimerInputs[1].Value()
				if timerName == "" || timerDuration == "" {
					return m, nil
				}

				timerDurationInt, _ := strconv.Atoi(timerDuration)
				m.timerManager.NewTimer(timerName, timerDurationInt)
				m.viewState = mainTimerPage
				m.addTimerInputs[m.focusIndex].Blur()
				m.focusIndex = 0
				return m, nil
			}
		}

		var formInputCmd tea.Cmd
		m.addTimerInputs[m.focusIndex], formInputCmd = m.addTimerInputs[m.focusIndex].Update(msg)
		batchedCmds = append(batchedCmds, formInputCmd)
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
