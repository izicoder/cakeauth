package main

import (
	"fmt"
	"izicoder/cakeauth/core"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println("starting")
	err := core.StartServer()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("stopping")
}
