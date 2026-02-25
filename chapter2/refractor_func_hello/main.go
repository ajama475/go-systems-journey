package main

import "fmt"

func main() {
	//calls on greet fun
	greeting := greet()
	fmt.Println(greeting)
}

// greet returns a greeting to the world
func greet() string {
	//return a simple greeting
	return "Hello, world"

}