package main

import (
	"fmt"
)

const (
	french     = "French"
	portuguese = "Portuguese"
	spanish    = "Spanish"

	englishHelloPrefix    = "Hello, "
	frenchHelloPrefix     = "Bonjour, "
	portugueseHelloPrefix = "Olá, "
	spanishHelloPrefix    = "Hola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
