package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/antoniofmoliveira/cleanarch/pkg/events"
	"github.com/streadway/amqp"
)

type OrderListedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderListedHandler(rabbitMQChannel *amqp.Channel) *OrderListedHandler {
	return &OrderListedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderListedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order listed: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
