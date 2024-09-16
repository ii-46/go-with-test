package main

import "fmt"

const lao = "Lao"
const french = "French"
const englishHelloPrefix = "Hello, "
const laoHelloPrefix = "ສະບາຍດີ, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		return englishHelloPrefix + "world"
	}
	return GreetingPrefix(lang) + name
}

func GreetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case lao:
		prefix = laoHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
