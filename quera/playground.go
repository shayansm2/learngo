package main

import (
	"errors"
	"fmt"
)

func main() {
	// Create a root error
	rootErr := errors.New("Root error")
	// Wrap the root error with additional context
	wrappedErr1 := fmt.Errorf("Error 1: %w", rootErr)
	wrappedErr2 := fmt.Errorf("Error 2: %w", wrappedErr1)
	fmt.Println(wrappedErr2)

	err1 := errors.New("error one")
	err2 := errors.New("error two")
	er3 := fmt.Errorf("error three")
	err := errors.Join(err1, err2, er3)
	fmt.Println(err)

	panic()
}
