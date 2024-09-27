package main

import (
	"fmt"
	"izicoder/cakeauth/core"
	"log"
)

func main() {
	fmt.Println("starting")
	// core.TestDB()
	log.Fatal(core.StartServer())
	fmt.Println("stopping")
}
