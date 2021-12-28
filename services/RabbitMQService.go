package services

import (
	"RestApiWithRedisAndRabbitMQ/config"
	"RestApiWithRedisAndRabbitMQ/models"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func ProduceMessages(student models.Student) int64{
	channel:= config.InitRabbitMQ() // create the postgres db connection
	var status int64  = 1
	json, _ := json.Marshal(student)
	err := channel.Publish(
		os.Getenv("RABBITMQ_EXCHANGE"),
		os.Getenv("RABBITMQ_ROUTE_KEY"),
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:        json,
		},
	)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into RabbitMQ", err)
		return 0
	}

	return status
}


func ConsumeMessages() ([]models.Student, int64) {
	var students []models.Student
	var status int64  = 1
	channel:= config.InitRabbitMQ() // create the postgres db connection
	messages, err := channel.Consume(
		os.Getenv("RABBITMQ_QUEUE"),
		os.Getenv("RABBITMQ_CONSUMER"),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into RabbitMQ", err)
		return students, 1
	}

	//forever := make(chan bool) // uncomment for using goroutine

	go func() {
	for message := range messages {
		var student models.Student
		json.Unmarshal(message.Body, &student)
		students = append(students, student)
		log.Printf("Received a message: %s", message.Body)
	}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	//<-forever // uncomment uncomment for using goroutine

	return students,status
}