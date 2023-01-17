package main

import (
	"context"
	"fmt"

	"github.com/bmviniciuss/golang-rest-api/internal/comment"
	"github.com/bmviniciuss/golang-rest-api/internal/db"
)

func Run() error {
	fmt.Println("Starting application...")
	db, err := db.NewDatabase()

	if err != nil {
		fmt.Println("Failed to connect to database", err)
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database", err)
		return err
	}

	fmt.Println("Connected to database successfully! :D")

	cmtService := comment.NewService(db)

	p, err := cmtService.CreateComment(context.Background(), comment.Comment{
		Slug:   "test",
		Body:   "test",
		Author: "Vinicius Barbosa",
	})

	if err != nil {
		return err
	}

	fmt.Println("Comment created: ", p)

	cmt, err := cmtService.GetComment(context.Background(), p.ID)

	if err != nil {
		fmt.Println("Failed to get comment", err)
		return err
	}

	fmt.Println("Comment: ", cmt)

	return nil
}

func main() {
	fmt.Println("Go Rest API")
	err := Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
