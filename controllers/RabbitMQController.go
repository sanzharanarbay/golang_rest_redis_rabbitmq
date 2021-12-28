package controllers

import (
	"RestApiWithRedisAndRabbitMQ/models"
	"RestApiWithRedisAndRabbitMQ/services"
	"encoding/json"
	"log"
	"net/http"
)

// Send Messages to RabbitMQ
func  SendRabbitMQMessages(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var student models.Student
	err:= json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call insert user function and pass the user
	insertID := services.ProduceMessages(student)

	// format a response object
	res := Response{
		Status:      insertID,
		Message: "Student sended to rabbitmq successFully!!!",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// Get Messages from RabbitMQ
func GetRabbitMQMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// call the getUser function with user id to retrieve a single user
	user, err := services.ConsumeMessages()
	if err == 0 {
		log.Fatalf("Unable to get user from Redis. %v", err)
	}
	// send the response
	json.NewEncoder(w).Encode(user)
}