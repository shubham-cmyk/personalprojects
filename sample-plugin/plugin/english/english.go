package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("Hello Universe")
}

func (g greeting) GoodBye() {
	fmt.Println("Goodbye")
}

// this is exported
var Greeter greeting
