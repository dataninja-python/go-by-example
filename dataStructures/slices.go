package main

import "fmt"

func main() {
	slice := make([]string, 3)
	fmt.Println("emp:", slice)

	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	fmt.Println("set: ", slice)
	fmt.Println("get: ", slice[2])

	fmt.Println("len: ", len(slice))

	slice = append(slice, "d")
	slice = append(slice, "e", "f")
	fmt.Println("apd: ", slice)

	sliceCopy := make([]string, len(slice))
	copy(sliceCopy, slice)
	fmt.Println("cpy: ", sliceCopy)

	l := slice[2:5]
	fmt.Println("sl1: ", l)

	l = slice[:5]
	fmt.Println("sl2: ", l)

	l = slice[2:]
	fmt.Println("sl3: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
		fmt.Println("2d: ", twoD)
	}
}
