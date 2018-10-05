package main

import "fmt"

func main() {
	var small int
	var big int

	fmt.Printf("Enter a number: ")
	fmt.Scan(&small)
	fmt.Printf("Enter a bigger number: ")
	fmt.Scan(&big)

	remainder := big % small
	fmt.Printf("%v divded by %v has a remainder of %v\n", big, small, remainder)
}
