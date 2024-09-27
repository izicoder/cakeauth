package main

import (
	"fmt"
	"izicoder/cakeauth/core"
)

func main() {
	fmt.Println("hello")
	fmt.Printf("1+2=%d", core.Add(1, 2))
	core.TestPG()
}
