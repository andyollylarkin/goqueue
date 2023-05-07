package eventbus

import (
	"github.com/sirupsen/logrus"
)

type EventBus struct {
	subscribers map[string][]*Subscriber // map[event type name] subscribers for this event type
	logger      *logrus.Logger
}

func NewEventBus(logger *logrus.Logger) *EventBus {
	return &EventBus{subscribers: make(map[string][]*Subscriber), logger: logger}
}

func (b *EventBus) Notify(event Event) {
	for k, v := range b.subscribers {
		if k == event.EventType() {
			for _, s := range v {
				(*s).OnEvent(event)
			}
		}
	}
}

func (b *EventBus) EmitEvent(event Event) {
	b.logger.WithFields(logrus.Fields{"event_type": event.EventType()}).Debug("Emitted new event")
	// count amount of subscribers for event
	for k, v := range b.subscribers {
		if k == event.EventType() {
			event.CountSubscribers(len(v))
			break
		}
	}
	b.Notify(event)
}

func (b *EventBus) RegisterSubscriber(eventType string, sub *Subscriber) {
	if s, ok := b.subscribers[eventType]; ok {
		s = append(s, sub)
	} else {
		b.subscribers[eventType] = []*Subscriber{sub}
	}
}

func (b *EventBus) DetachSubscriber(eventName string, sub *Subscriber) {
	if s, ok := b.subscribers[eventName]; ok {
		for i, v := range s {
			if v == sub {
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
}
