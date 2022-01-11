package main

import (
	"editt/database"
	"editt/routers"
	"log"
)

func main() {
	database.Setup()
	r := routers.Setup()
	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}