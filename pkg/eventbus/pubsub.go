package eventbus

type Subscriber interface {
	OnEvent(event Event)
}
