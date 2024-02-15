package main

import (
	"context"
	"fmt"
	"log"

	"github.com/roblesdotdev/go-rest-api/internal/comment"
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
	}

	if err := db.MigrateDB(); err != nil {
		log.Println("failed to migrate database")
		return err
	}

	log.Println("Successful db connection")
	cmt := comment.Comment{
		Slug:   "new-comment",
		Body:   "Hello, World",
		Author: "Robles AR",
	}
	fmt.Println(db.CreateComment(context.Background(), cmt))

	return nil
}

func main() {
	log.Println("Go REST API")
	if err := Run(); err != nil {
		log.Fatalf("Fail to start app %V\n", err)
	}
}
