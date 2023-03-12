package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("你好宇宙")
}

func (g greeting) GoodBye() {
	fmt.Println("再见")
}

// this is exported
var Greeter greeting
