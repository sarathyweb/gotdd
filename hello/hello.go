package main

import "fmt"

const (
	englishHelloPrefix  = "Hello, "
	spainishHelloPrefix = "Hola, "
	frenchHelloPrefix   = "Bonjour, "

	spainish = "Spainish"
	french   = "French"
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
	case spainish:
		prefix = spainishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
