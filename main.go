package main

import (
	"sync"
	"time"

	linkedlist "github.com/hmcalister/Go-DSA/list/LinkedList"
	"github.com/hmcalister/TeaTimer/internal/timerdata"
)

func main() {

	var timerListMutex sync.RWMutex
	allTimers := linkedlist.New[*timerdata.TimerData]()
	globalTimer := time.NewTicker(time.Second)
	go func() {
		for range globalTimer.C {
			timerListMutex.RLock()
			linkedlist.ForwardApply(allTimers, func(item *timerdata.TimerData) {
				item.UpdateChannel <- timerdata.UpdateMessageTick
			})
			timerListMutex.RUnlock()
		}
	}()

	t1 := timerdata.NewTimer("Timer A", 7)
	defer close(t1.UpdateChannel)
	allTimers.Add(t1)

	t2 := timerdata.NewTimer("Timer B", 7)
	defer close(t2.UpdateChannel)
	allTimers.Add(t2)

	time.Sleep(3 * time.Second)
	timerListMutex.Lock()
	close(t1.UpdateChannel)
	allTimers.RemoveAtIndex(0)
	timerListMutex.Unlock()
	for {
	}
}
