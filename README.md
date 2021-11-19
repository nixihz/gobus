# Gobus 🚧
The project is under development, please wait.

# Usage

```go

	demoUpdated := demo.NewUpdatedEvent("8002", "Order Updated")

	// local adapter
	publisher := gobus.GetEventPublisher(impl.NewLocalAdapter())
	publisher.PublishEvent(gobus.Event(demoUpdated))

	// rabbitmq adapter
	amqpURI := "amqp://username:password@host:port/"
	publisher2 := gobus.GetEventPublisher(impl.NewRabbitMqAdapter(amqpURI))
	publisher2.PublishEvent(gobus.Event(demoUpdated))

```
