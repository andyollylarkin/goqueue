package eventbus

type Event interface {
	// EventType returns string representation of event type
	EventType() string
	// ProcessEvent do processing event
	ProcessEvent(processHandler func(opts ...any), opts ...any)
	// CountSubscribers should increment amount of subscribers.
	//Amount of subscribers use for wait until each of them done work
	CountSubscribers(numberOfSubscribers int)
	WaitUntilEventWillBeProcessed()
}
