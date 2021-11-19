package impl

import "github.com/nixihz/gobus"

func NewLocalAdapter() *LocalAdapter {
	r := LocalAdapter{
		BusName:  "gobus",
		Exchange: "gobus-exchange",
	}

	return &r
}

type LocalAdapter struct {
	gobus.EventPublisher
	BusName  string
	Exchange string
}

func (adapter *LocalAdapter) PublishEvent(event gobus.Event) error {
	println(string(event.Serialize()))

	return nil
}
