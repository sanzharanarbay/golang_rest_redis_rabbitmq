package controllers

import (
	"RestApiWithRedisAndRabbitMQ/models"
	"RestApiWithRedisAndRabbitMQ/services"
	"encoding/json"
	"log"
	"net/http"
)

// Set key to Redis
func  InsertRedisUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err:= json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call insert user function and pass the user
	insertID := services.SetRedis(user)

	// format a response object
	res := Response{
		Status:      insertID,
		Message: "User inserted to redis successFully!!!",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// Get value by key in Redis
func GetRedisUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get the userid from the request params, key is "id"
	// convert the id type from string to int
	iin:= r.FormValue("iin")
	// call the getUser function with user id to retrieve a single user
	user, err := services.GetRedis(iin)
	if err == 0 {
		log.Fatalf("Unable to get user from Redis. %v", err)
	}
	// send the response
	json.NewEncoder(w).Encode(user)
}

// Delete value by key in Redis
func DeleteRedisUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get the userid from the request params, key is "id"
	// convert the id type from string to int
	iin:= r.FormValue("iin")
	// call the getUser function with user id to retrieve a single user
	id := services.DeleteRedis(iin)
	res := Response{
		Status:      id,
		Message: "User inserted to redis successFully!!!",
	}
	// send the response
	json.NewEncoder(w).Encode(res)
}

