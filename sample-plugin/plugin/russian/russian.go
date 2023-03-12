package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("привет вселенная")
}
func (g greeting) GoodBye() {
	fmt.Println("до свидания")
}

// this is exported
var Greeter greeting
