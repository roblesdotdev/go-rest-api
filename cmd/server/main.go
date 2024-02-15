package main

import (
	"context"
	"log"

	"github.com/roblesdotdev/go-rest-api/internal/comment"
	"github.com/roblesdotdev/go-rest-api/internal/db"
	transportHttp "github.com/roblesdotdev/go-rest-api/internal/transport/http"
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

	cmtService := comment.NewService(db)
	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	log.Println("Go REST API")
	if err := Run(); err != nil {
		log.Fatalf("Fail to start app %V\n", err)
	}
}
