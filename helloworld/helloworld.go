package main

const helloPreFix = "hello, "
const holaPreFix = "hola, "
const englishPrefix = "Hello, "
const spanish = "spanish"
const french = "french"

// Hi should return hello, <name> in <language>
func Hi(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = "hola, "
	case french:
		prefix = "bonjour, "
	default:
		prefix = "hello, "
	}
	return
}
