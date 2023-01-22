package core

import (
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
