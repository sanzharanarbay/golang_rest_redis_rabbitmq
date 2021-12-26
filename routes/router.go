package routes

import (
	"RestApiWithRedisAndRabbitMQ/controllers"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/orders/order/{id}", controllers.GetOrder).Methods("GET")
	router.HandleFunc("/api/orders/all", controllers.GetOrders).Methods("GET")
	router.HandleFunc("/api/orders/create", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/api/orders/update/{id}", controllers.UpdateOrder).Methods("PUT")
	router.HandleFunc("/api/orders/delete/{id}", controllers.DeleteOrder).Methods("DELETE")
	return router
}
