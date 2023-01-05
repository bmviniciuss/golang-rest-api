package main

import (
	"context"
	"fmt"

	"github.com/bmviniciuss/golang-rest-api/internal/db"
)

func Run() error {
	fmt.Println("Starting application...")
	db, err := db.NewDatabase()

	if err != nil {
		fmt.Println("Failed to connect to database", err)
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("Connected to database successfully! :D")

	return nil
}

func main() {
	fmt.Println("Go Rest API")
	err := Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
