package main

import(
	"fmt"
)

func main() {
	greeting := greet("hi")
	fmt.Println(greeting)
}

type language string

var phrasebook= map[language]string{
	"en": "Hello, world",		 // English
	"fr": "Bonjour le monde",	 // French
	"ar": "مرحبا بالعالم" ,		 // Arabic
	"ch": "你好世界",			  // Chinese Simplified
	"hi": "हैलो वर्ल्ड",				 // Hindi
}

func greet(l language) string {  
	greeting, ok := phrasebook[l]

	if !ok {
		return fmt.Sprintf("Unsupported language: %q", l)
	}

	return greeting

}