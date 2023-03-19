package app

type ShutdownEvent struct {
	amountOfSubscribers int
	subscriberDoneChan  chan struct{}
}

const ShutdownEventType = "shutdown_event"

func (s *ShutdownEvent) EventType() string {
	return ShutdownEventType
}

func (s *ShutdownEvent) ProcessEvent(processHandler func(opts any), opts any) {
	processHandler(opts)
	s.subscriberDoneChan <- struct{}{}
}

func (s *ShutdownEvent) WaitUntilEventWillBeProcessed() {
	for i := 0; i < s.amountOfSubscribers; i++ {
		<-s.subscriberDoneChan
	}
}

func (s *ShutdownEvent) RegisterSubscriber(numberOfSubscribers int) {
	s.amountOfSubscribers = numberOfSubscribers
}

func NewShutdownEvent() *ShutdownEvent {
	return &ShutdownEvent{}
}
