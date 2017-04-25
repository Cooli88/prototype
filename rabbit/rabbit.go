package rabbit

import (
	"github.com/streadway/amqp"
	"log"
)

const (
	rabbitUrlConnect = "amqp://rabbit:ujkjdjyju@parking-rabbit-dev"
	queueName        = "prkm-device-message"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//refactor
func SendMessage(body []byte) {
	conn, err := amqp.Dial(rabbitUrlConnect)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
