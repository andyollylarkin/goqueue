package broker

import (
	"goqueue/pkg/eventbus"
)

type Broker struct {
	events chan eventbus.Event
}

func NewBroker(events chan eventbus.Event) *Broker {
	return &Broker{events: events}
}
