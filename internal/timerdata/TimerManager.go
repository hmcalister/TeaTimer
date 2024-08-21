package timerdata

import (
	"sync"
	"time"

	linkedlist "github.com/hmcalister/Go-DSA/list/LinkedList"
)

type TimerManager struct {
	allTimersMutex sync.RWMutex
	allTimers      *linkedlist.LinkedList[*TimerData]

	globalTicker *time.Ticker
}

func NewManager() *TimerManager {
	manager := &TimerManager{
		allTimers:    linkedlist.New[*TimerData](),
		globalTicker: time.NewTicker(time.Second),
	}

	go func() {
		for range manager.globalTicker.C {
			manager.allTimersMutex.RLock()
			go linkedlist.ForwardApply(manager.allTimers, func(item *TimerData) {
				item.UpdateChannel <- UpdateMessageTick
			})
			manager.allTimersMutex.RUnlock()
		}
	}()

	return manager
}

// Creates and returns a new Timer.
//
// When finished with the timer, close the update channel to signal cleanup.
func (manager *TimerManager) NewTimer(name string, duration int) {
	updateChannel := make(chan TimerUpdateMessageEnum)
	t := &TimerData{
		Name:              name,
		TimerState:        TimerStateRunning,
		InitialDuration:   duration,
		RemainingDuration: duration,
		UpdateChannel:     updateChannel,
	}

	started := make(chan interface{})
	go t.stateMachine(started)
	<-started

	// Now timer has started processing events, we can add it to the timer list

	manager.allTimersMutex.Lock()
	manager.allTimers.Add(t)
	manager.allTimersMutex.Unlock()
}
