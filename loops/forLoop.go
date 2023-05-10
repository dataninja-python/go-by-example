// practice for loops

package main

import "fmt"

func main() {
	// starting from 1 loop until end of loop
	i := 1
	end := 3
	for i <= end {
		if i == 1 {
			fmt.Println("This is the first time through the loop")
		}
		if i == end {
			fmt.Println("This is the last time through the loop")
		}
		fmt.Println("Current value is: ", i) // print the current value
		i = i + 1
	}
}
