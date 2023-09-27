package main

import "fmt"

func main() {
	//This is a comment
	var name string

	fmt.Println("what is your name?")

	fmt.Scan(&name)

	fmt.Println("Hello", name)
}
