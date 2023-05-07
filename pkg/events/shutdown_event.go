package events

import "time"

type ShutdownEvent struct {
	amountOfSubscribers int
	subscriberDoneChan  chan struct{}
}

const ShutdownEventType = "shutdown_event"

func (s *ShutdownEvent) EventType() string {
	return ShutdownEventType
}

func (s *ShutdownEvent) ProcessEvent(processHandler func(opts ...any), opts ...any) {
	processHandler(opts)
	s.subscriberDoneChan <- struct{}{}
}

func (s *ShutdownEvent) WaitUntilEventWillBeProcessed() {
	for i := 0; i < s.amountOfSubscribers; i++ {
		select {
		case <-s.subscriberDoneChan:
		//if the subscribers were lost or for some reason cannot complete the work after 3 minutes, we forcefully close the work
		case <-time.After(time.Minute * 3):
			return
		}
	}
}

func (s *ShutdownEvent) CountSubscribers(numberOfSubscribers int) {
	s.amountOfSubscribers = numberOfSubscribers
}

func NewShutdownEvent() *ShutdownEvent {
	return &ShutdownEvent{amountOfSubscribers: 0, subscriberDoneChan: make(chan struct{}, 0)}
}
