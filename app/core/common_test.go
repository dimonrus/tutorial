package core

import (
	"context"
	"github.com/dimonrus/gocli"
	"github.com/dimonrus/gohelp"
	amqp "github.com/rabbitmq/amqp091-go"
	"testing"
	"tutorial/app/base"
)

func TestCommon(t *testing.T) {
	base.App.SuccessMessage("Tests are work's fine")
}

func TestDatabaseVersion(t *testing.T) {
	var version string
	db := base.App.GetDB()
	err := db.QueryRow("SELECT VERSION();").Scan(&version)
	if err != nil {
		t.Fatal(err)
	}
	base.App.SuccessMessage("Postgres version is: " + version)
}

func TestRabbitMQ(t *testing.T) {
	var message = []byte("some data")
	p := amqp.Publishing{Body: message}
	broker := base.App.GetRabbit()
	e := broker.Publish(p, "tutorial.test", "local")
	if e != nil {
		t.Fatal(e)
	}
	base.App.SuccessMessage("Publish message into queue ")
}

func TestContextLoggerOutput(t *testing.T) {
	t.Run("info", func(t *testing.T) {
		config := gocli.LoggerConfig{Level: gocli.LogLevelDebug}
		logger := gocli.NewLogger(config)
		logger.Infoln("Some info")
	})
	t.Run("logger_forma", func(t *testing.T) {
		config := gocli.LoggerConfig{
			Level:  gocli.LogLevelErr,
			Format: gocli.LoggerFormat{"application": "app: %s ", "requestId": "rid: %s"},
		}
		ctx := context.WithValue(context.Background(), "application", "tutorial")
		ctx = context.WithValue(ctx, "requestId", gohelp.NewUUID())
		logger := gocli.NewLogger(config)
		// ignored according level
		logger.Info("ignored")
		// output with info from context
		logger.Errorln(ctx, "Error happened")
	})
}
