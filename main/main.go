package main

import (
	"github.com/nixihz/gobus"
	"github.com/nixihz/gobus/demo"
	"github.com/nixihz/gobus/impl"
)

func main() {
	demoUpdated := demo.NewUpdatedEvent("8002", "Order Updated")

	// local adapter
	publisher := gobus.GetEventPublisher(impl.NewLocalAdapter())
	publisher.PublishEvent(gobus.Event(demoUpdated))

	// rabbitmq adapter
	amqpURI := "amqp://username:password@host:port/"
	publisher2 := gobus.GetEventPublisher(impl.NewRabbitMqAdapter(amqpURI))
	publisher2.PublishEvent(gobus.Event(demoUpdated))

}
