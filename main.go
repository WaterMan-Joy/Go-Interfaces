package main

import "fmt"

type bot interface {
	getGreeting() string
}

type engBot struct{}
type spaBot struct{}

func main() {
	eb := engBot{}
	sb := spaBot{}
	printFunc(eb)
	printFunc(sb)
}

func printFunc(b bot) {
	fmt.Println(b.getGreeting())
}

func (engBot) getGreeting() string {
	return "Hello, I am an English bot "
}

func (spaBot) getGreeting() string {
	return "Hello, I am an English bot "
}
