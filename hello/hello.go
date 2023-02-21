package main

import (
	"fmt"

	"github.com/dandawong/go-learning/greetings"
)

func main() {
	fmt.Println("Hello, World!")
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
