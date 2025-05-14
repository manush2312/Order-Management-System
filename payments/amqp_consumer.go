package main

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/manush2312/commons/api"
	"github.com/manush2312/commons/broker"
	amqp "github.com/rabbitmq/amqp091-go"
)

type consumer struct {
	service PaymentsService
}

func NewConsumer(service PaymentsService) *consumer {
	return &consumer{service: service}
}

func (c *consumer) Listen(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(broker.OrderCreatedEvent, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Recieved message: %s", d.Body)

			o := &pb.Order{}
			if err := json.Unmarshal(d.Body, o); err != nil {
				log.Printf("failed to unmarshal order %v", err)
				continue
			}
			paymentLink, err := c.service.CreatePayment(context.Background(), o)
			if err != nil {
				log.Printf("failed to create payment %v", err)
				continue
			}
			log.Printf("payment link created %s", paymentLink)
		}
	}()
	<-forever
}
