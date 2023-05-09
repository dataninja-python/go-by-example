package main

import (
	"fmt"
	"math"
)

// assign a constant variable
const s string = "constant"

func main() {
	fmt.Println("const s string = 'constant' = ", s)
	`
	const n = 500000000`

	const d = 3e20 / n
	fmt.Println("const n = ", n)
	fmt.Println("3e20 = ", 3e20)
	fmt.Println("const d = 3e20 / n", d)

	fmt.Println("int64(d) = ", int64(d))

	fmt.Println("math.Sin(n) = ", math.Sin(n))
}
