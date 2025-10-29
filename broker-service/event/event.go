package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declerExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", //name
		"topic",      //type
		true,         //durable?
		false,        //auto-deletad
		false,        //internal
		false,        //no wait
		nil,          //arguments
	)
}

func declerRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
}
