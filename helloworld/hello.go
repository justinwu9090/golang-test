// package to group up my code
package main

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// function Hello returns a string
func Hello(name string, language string) string { // parameters are name type instead of type name
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)
	return prefix + name // no semicolon
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

// main function called in start
// func main() {
// 	fmt.Println(Hello())
// }
