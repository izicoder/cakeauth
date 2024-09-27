package core

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func Add(a int, b int) int {
	ptr := jwt.New(jwt.SigningMethodES256)
	what := ptr.Header
	fmt.Printf("what: %v\n", what)
	return a + b
}
