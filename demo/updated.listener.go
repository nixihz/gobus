package demo

import "github.com/nixihz/gobus"

type UpdatedListener struct {
	gobus.Listener
}

func (updatedListener *UpdatedListener) OnEvent(event UpdatedEvent) {
	println("Updated Received!")
}
