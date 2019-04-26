package main

import "fmt"

func Hello() string {
	return "hello, world"
}

func Hi(name string) string {
	return "hello, " + name
}

func main() {
	fmt.Println(Hello())
	fmt.Println(Hi("ben"))
}
