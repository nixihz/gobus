package impl

import (
	"fmt"
	gobus "github.com/nixihz/gobus"
	"github.com/streadway/amqp"
)

func NewRabbitMqAdapter(amqpURI string) *RabbitMqAdapter {
	r := RabbitMqAdapter{
		AmqpUri:  amqpURI,
		BusName:  "gobus",
		Exchange: "gobus-exchange",
	}

	return &r
}

type RabbitMqAdapter struct {
	gobus.EventPublisher
	AmqpUri  string
	BusName  string
	Exchange string
}

func (adapter *RabbitMqAdapter) PublishEvent(event gobus.Event) error {
	conn := NewConnection(
		adapter.AmqpUri,
		adapter.BusName,
		adapter.Exchange,
		[]string{"queue-1", "queue-2"},
	)
	if err := conn.Connect(); err != nil {
		panic(err)
	}
	if err := conn.BindQueue(); err != nil {
		panic(err)
	}
	select { //non blocking channel - if there is no error will go to default where we do nothing
	case err := <-conn.err:
		if err != nil {
			conn.Reconnect()
		}
	default:
	}

	p := amqp.Publishing{
		ContentType: "text/plain",
		Body:        event.Serialize(),
	}
	if err := conn.channel.Publish(conn.exchange, "", false, false, p); err != nil {
		return fmt.Errorf("Error in Publishing: %s", err)
	}
	return nil
}
