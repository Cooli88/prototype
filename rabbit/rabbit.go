package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

const (
	rabbitUrlConnect  = "amqp://rabbit:ujkjdjyju@parking-rabbit-dev"
	rabbitUrlTemplate = "amqp://%s:%s@%s"
	queueName         = "prkm-device-message"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//refactor
func SendMessage(body []byte) {
	urlToConnectRabbit := fmt.Sprintf(rabbitUrlTemplate, os.Getenv("RABBIT_USER"), os.Getenv("RABBIT_PASSWORD"), os.Getenv("RABBIT_HOST"))
	conn, err := amqp.Dial(urlToConnectRabbit)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
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
