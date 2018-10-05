package main

import "fmt"

func main() {
	var name string
	fmt.Printf("Type in your name:")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
}
