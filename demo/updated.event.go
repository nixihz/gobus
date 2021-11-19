package demo

import "github.com/nixihz/gobus"

type UpdatedEvent struct {
	Code string
	Msg  string
}

func NewUpdatedEvent(code string, msg string) *gobus.ApplicationEvent {
	return &gobus.ApplicationEvent{
		Name: "",
		Content: UpdatedEvent{
			Code: code,
			Msg:  msg,
		},
	}
}
