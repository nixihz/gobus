package gobus

type Listener interface {
	OnEvent(event Event)
}
