package config

import (
	"fmt"
	"github.com/joho/godotenv" // package used to read the .env file
	"github.com/streadway/amqp"
	"log"
	"os" // used to read the environment variable
)

func InitRabbitMQ() *amqp.Channel {
	e:= godotenv.Load()
	if e != nil {
		log.Fatalf("Error loading .env file")
	}

	rabbitmq_host := os.Getenv("RABBITMQ_HOST")
	rabbitmq_vhost := os.Getenv("RABBITMQ_VHOST")
	rabbitmq_login := os.Getenv("RABBITMQ_LOGIN")
	rabbitmq_password := os.Getenv("RABBITMQ_PASSWORD")
	rabbitmq_queue := os.Getenv("RABBITMQ_QUEUE")
	rabbitmq_exchange := os.Getenv("RABBITMQ_EXCHANGE")
	rabbitmq_key := os.Getenv("RABBITMQ_ROUTE_KEY")

	rabbitUri := fmt.Sprintf("amqp://%s:%s@%s%s", rabbitmq_login, rabbitmq_password, rabbitmq_host, rabbitmq_vhost) //Build connection string

	conn, err := amqp.Dial(rabbitUri)
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	//defer ch.Close()

	_, err = ch.QueueDeclare(
		rabbitmq_queue,    // name
		true, // durable
		false, // delete when unused
		false,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	err = ch.ExchangeDeclare(
		rabbitmq_exchange, // name
		"direct",           // type
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)

	err = ch.QueueBind(
		rabbitmq_queue, // queue name
		rabbitmq_key,     // routing key
		rabbitmq_exchange, // exchange
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	return ch
}
