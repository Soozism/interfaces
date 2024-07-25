package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface{
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(sb)
	printGreeting(eb)

}

func (englishBot) getGreeting() string {
	return "hi"
}

func (spanishBot) getGreeting() string {
	return "hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

