// Practice using if and if/else states to make decisions in Go

package main

import "fmt"

func main() {
	// use the modulus operator to determine if a number is even or odd
	numbers := []int{1, 12, 33, 4, 59, 61, 7, 8, 16, 10, 11, 20, 13, 14, 15, 16, 17, 18, 19}
	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Println(num, " is even")
		} else {
			fmt.Println(num, " is odd")
		}
	}
	fmt.Println("--------------------")
	// use the modulus operator to determine if a number is divisible by 4
	for _, num := range numbers {
		if num%4 == 0 {
			fmt.Println(num, " is divisible by 4")
		}
	}
	fmt.Println("--------------------")
	// use the modulus operator to determine if a number is negative or has multiple digits
	for _, num := range numbers {
		if num < 0 { // negative numbers are less than 0 on the number line
			fmt.Println(num, " is negative")
		} else if num < 10 { // numbers less than 10 have only 1 digit
			fmt.Println(num, " has 1 digit")
		} else { // numbers greater than 10 have multiple digits
			fmt.Println(num, " has multiple digits")
		}
	}
	fmt.Println("--------------------")
	fmt.Println("NOTE: Go, like Haskell does not have a ternary operator like Python, Javascript, Lisp, " +
		"C++ and many other languages")
}
