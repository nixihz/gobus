package gobus

import (
	"encoding/json"
	"fmt"
)

type EventLabel string

type Event interface {
	Label() EventLabel
	Serialize() []byte
}

type Content interface {
}

type ApplicationEvent struct {
	Name    string
	Content Content
}

func (appEvent *ApplicationEvent) Label() EventLabel {
	return EventLabel(appEvent.Name)
}

func (appEvent *ApplicationEvent) Serialize() []byte {
	buf, err := json.Marshal(appEvent)
	if err != nil {
		fmt.Errorf("marshal value for event %s: %w", appEvent, err)
	}
	return buf
}

type EventPublisher interface {
	PublishEvent(event Event) error
}

func GetEventPublisher(publisher EventPublisher) EventPublisher {
	return publisher
}
