package main

import (
	"log"
)

// Run - is going to be responsible for
// the instantiation and startup of our
// go application.
func Run() error {
	log.Println("Starting up our application")
	return nil
}

func main() {
	log.Println("Go REST API")
	if err := Run(); err != nil {
		log.Fatalf("Fail to start app %V\n", err)
	}
}
