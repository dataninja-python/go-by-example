package main

import "fmt"

func main() {
	// declare a variable
	var a = "initial"
	fmt.Println("var a = ", a)

	// declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println("var b, c = ", b, c)

	// infer type
	var d = true
	fmt.Println("var d = ", d)

	// declare without initializing
	var e int
	fmt.Println("var e = ", e)

	// shorthand for declaring and initializing
	f := "short"
	fmt.Println("f := 'short' = ", f)
}
