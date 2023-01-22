// event consumer file
package consumer

import (
	"github.com/dimonrus/gorabbit"
	amqp "github.com/rabbitmq/amqp091-go"
	"tutorial/app/base"
)

func init() {
	base.App.GetRabbit().GetRegistry()["event"] = &gorabbit.Consumer{Queue: "tutorial.event", Server: "local", Count: 1, Callback: func(d amqp.Delivery) {
		base.App.SuccessMessage("We have a new event: " + string(d.Body))
	}}
}
