package main

import "fmt"

func Run() error {
	fmt.Println("Running...")
	return nil
}

func main() {
	fmt.Println("Go Rest API")
	err := Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
