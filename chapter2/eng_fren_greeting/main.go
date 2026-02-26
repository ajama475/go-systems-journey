package main

import "fmt"

func main() {
	greeting := greet("fr")
	fmt.Println(greeting)
}

type language string

func greet(lan language) string {
	switch lan {
	case "en":
		return "Hello, world"
	case "fr":
		return "Bonjour le monde"
	default:
		return ""
	}
}