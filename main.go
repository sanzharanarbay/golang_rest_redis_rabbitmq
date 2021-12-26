package main

import(
	"RestApiWithRedisAndRabbitMQ/routes"
	"fmt"
	"net/http"
)

func main() {
	router := routes.Router()
	fmt.Println("Starting server on the port 8080...")
	http.ListenAndServe(":8080", router)
}
