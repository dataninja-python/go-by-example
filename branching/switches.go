// special kind of conditional that tests statements across many cases simultaneously
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Switches in Go are a convenient way to express a multi-way if-else statement.")
	fmt.Println("They are used to test many conditions (called cases) simultaneously.")

	fmt.Println("-----------------------------")
	i := 2
	fmt.Println("i := 2 = ", i)
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("ERROR: unknown number") // should not reach this case but always included for safety
	}

	fmt.Println("-----------------------------")
	fmt.Println("You can use Go's built-in time functions with switches to determine weekday or weekend.")
	currentDay := time.Now().Weekday()
	fmt.Println("Today is ", currentDay)
	switch currentDay {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday. Get to work!")
	}

	fmt.Println("-----------------------------")
	t := time.Now()
	fmt.Println("t := time.Now() = ", t)
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon.")
	default:
		fmt.Println("It's after noon.")
	}

	fmt.Println("-----------------------------")
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("I don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
