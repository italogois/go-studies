package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"
const french = "French"
const frenchHelloPrefix = "Bonjour, "
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	if language == spanish {
		prefix = spanishHelloPrefix
	}

	if language == french {
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Man", "Spanish"))
}

