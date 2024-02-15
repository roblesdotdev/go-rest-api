package main

import (
	"context"
	"log"

	"github.com/roblesdotdev/go-rest-api/internal/db"
)

// Run - is going to be responsible for
// the instantiation and startup of our
// go application.
func Run() error {
	log.Println("Starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		log.Println("Failed to connect to the database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	} else {
		log.Println("Successful db connection")
	}
	return nil
}

func main() {
	log.Println("Go REST API")
	if err := Run(); err != nil {
		log.Fatalf("Fail to start app %V\n", err)
	}
}
