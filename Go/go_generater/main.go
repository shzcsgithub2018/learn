package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("test")
	fmt.Printf("%s", err)
}
