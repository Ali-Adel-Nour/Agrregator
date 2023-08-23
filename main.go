package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Print("Hello world!")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the environment variable")
	}
	fmt.Println("Port: ", portString)
}
