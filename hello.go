package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const yoruba = "Yoruba"
const yorubaHelloPrefix = "Bawoni, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) { // named returned value, prefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case yoruba:
		prefix = yorubaHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main()  {
	fmt.Println(Hello("dre", ""))
}