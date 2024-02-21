package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("RabbitMQ in Golang: Getting started tutorial")

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	fmt.Println("Successfully connected to RabbitMQ instance")

	// opening a channel over the connection established to interact with RabbitMQ
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	// declaring queue with its properties over the the channel opened
	queue, err := channel.QueueDeclare(
		"test.direct", // name
		false,         // durable
		false,         // auto delete
		false,         // exclusive
		false,         // no wait
		nil,           // args
	)

	if err != nil {
		panic(err)
	}

	// publishing a message
	err = channel.Publish(
		"",            // exchange
		"test.direct", // key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Test Message 4"),
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Queue status:", queue)
	fmt.Println("Successfully published message")

}
