package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	prefix = englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return
	// here making prefix variable in the return statement, makes it such that new prefix variable is created
	// and writing a simple return just returns the prefix variable
}

// func Hello(name string, language string) string {
// 	// return "Hello, World!"

// 	if name == "" {
// 		name = "World"
// 	}

// 	// if language == spanish {
// 	// 	return spanishHelloPrefix + name + "!"
// 	// }

// 	// if language == french {
// 	// 	return frenchHelloPrefix + name + "!"
// 	// }

// 	prefix := englishHelloPrefix
// 	switch language {
// 	case spanish:
// 		prefix = spanishHelloPrefix
// 	case french:
// 		prefix = frenchHelloPrefix
// 	}

// 	return prefix + name + "!"
// }

// func HelloEnglish(name string) string {
// 	// return englishHelloPrefix + name + "!"
// 	if name == "" {
// 		name = "World"
// 	}
// 	return englishHelloPrefix + name + "!"

// }

func main() {
	// fmt.Println(Hello())
	// fmt.Println(HelloEnglish("Neel"))

	fmt.Println(Hello("Neel", ""))
}
