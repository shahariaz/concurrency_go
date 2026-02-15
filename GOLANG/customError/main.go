package main

import "fmt"

type CustomError struct {
	Message string
	Code    int
	err     error
}

func (e *CustomError) Error() string {
	fmt.Printf("Error occurred: %s (code: %d)\n", e.Message, e.Code)
	return fmt.Sprintf("Error %d: %s,", e.Code, e.Message)
}

func PrintName() error {
	return &CustomError{
		Message: "Name is empty",
		Code:    400,
	}
}

func main() {
	if err := PrintName(); err != nil {
		fmt.Println("An error occurred:", err.Error())
		return
	}
}
