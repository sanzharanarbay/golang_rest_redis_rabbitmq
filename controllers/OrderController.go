package controllers

import (
	"RestApiWithRedisAndRabbitMQ/models"
	"RestApiWithRedisAndRabbitMQ/services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// response format
type Response struct {
	Status      int64  `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

func  CreateOrder(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var order models.Order
	err:= json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call insert user function and pass the user
	insertID := services.InsertOrder(order)

	// format a response object
	res := Response{
		Status:      insertID,
		Message: "Order created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// GetUser will return a single order by its id
func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)
	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	// call the getUser function with user id to retrieve a single user
	user, err := services.GetOrder(int64(id))
	if err != nil {
		log.Fatalf("Unable to get order. %v", err)
	}
	// send the response
	json.NewEncoder(w).Encode(user)
}

// GetOrders will return all the users
func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get all the users in the db
	users, err := services.GetAllOrders()
	if err != nil {
		log.Fatalf("Unable to get all order. %v", err)
	}
	// send all the users as response
	json.NewEncoder(w).Encode(users)
}

// UpdateUser update user's detail in the postgres db
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)
	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	// create an empty user of type models.User
	var order models.Order
	// decode the json request to user
	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	// call update user to update the user
	updatedRows := services.UpdateOrder(int64(id), order)
	// format the message string
	msg := fmt.Sprintf("Order updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		Status:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// DeleteUser delete user's detail in the postgres db
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)
	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	// call the deleteUser, convert the int to int64
	deletedRows := services.DeleteOrder(int64(id))
	// format the message string
	msg := fmt.Sprintf("Order updated successfully. Total rows/record affected %v", deletedRows)
	// format the reponse message
	res := Response{
		Status:      int64(id),
		Message: msg,
	}
	// send the response
	json.NewEncoder(w).Encode(res)
}
