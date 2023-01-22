package consumer

import (
	"github.com/dimonrus/gorabbit"
	amqp "github.com/rabbitmq/amqp091-go"
	"tutorial/app/base"
)

func init() {
	base.App.GetRabbit().GetRegistry()["test"] = &gorabbit.Consumer{Queue: "tutorial.test", Server: "local", Count: 1, Callback: func(d amqp.Delivery) {
		// show the message
		base.App.GetLogger().Infof("%s", d.Body)
	}}
}
