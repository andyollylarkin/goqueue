package broker

import (
	"goqueue/pkg/eventbus"
)

type Topic struct {
	onEventCh chan eventbus.Event
}

func (t *Topic) OnEvent(event eventbus.Event) {
	t.onEventCh <- event
}
