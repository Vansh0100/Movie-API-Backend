package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Vansh0100/movieapi/controller"
	"github.com/Vansh0100/movieapi/router"
	"github.com/joho/godotenv"
)

// const connectionString = os.Getenv(mongoUrl)

func main() {
	// connection.Connection(connectionString)
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	connectionString := os.Getenv("MONGO_URL")
	controller.Connection(connectionString)
	r := router.Router()
	fmt.Println("Server is up and running....")
	log.Fatal(http.ListenAndServe(":4000", r))

}
