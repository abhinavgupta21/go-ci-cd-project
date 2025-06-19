package main

import "fmt"

func getGreeting() string {
	return "Hello, CI/CD with Github Actions Project!"
}

func main() {
	fmt.Println(getGreeting())
}
