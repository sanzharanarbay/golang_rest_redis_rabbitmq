package routes

import (
	"RestApiWithRedisAndRabbitMQ/controllers"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()
	// REST API
	router.HandleFunc("/api/orders/order/{id}", controllers.GetOrder).Methods("GET")
	router.HandleFunc("/api/orders/all", controllers.GetOrders).Methods("GET")
	router.HandleFunc("/api/orders/create", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/api/orders/update/{id}", controllers.UpdateOrder).Methods("PUT")
	router.HandleFunc("/api/orders/delete/{id}", controllers.DeleteOrder).Methods("DELETE")

	// REDIS API
	router.HandleFunc("/api/redis/set", controllers.InsertRedisUser).Methods("POST")
	router.Path("/api/redis/get").Queries("iin", "{[0-9]*?}").HandlerFunc(controllers.GetRedisUser).Methods("GET")
	router.Path("/api/redis/delete").Queries("iin", "{[0-9]*?}").HandlerFunc(controllers.DeleteRedisUser).Methods("GET")

	//RABBIT MQ API

	return router
}
