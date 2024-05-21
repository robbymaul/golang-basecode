package main

import (
	"basecode/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	route := routes.RouteInit()

	route.Run(":8008")
}
