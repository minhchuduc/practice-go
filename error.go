package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() (int, error) {
	return 5, &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main(a) {
	if a, err := run(); err != nil {
		fmt.Printf("a = %v but with error: %v", a, err)
	}
}
