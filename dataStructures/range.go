package main

import "fmt"

func main() {

	// use just value from range
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// use index (key) and value from range
	for i, num := range nums {
		if num == 3 {
			fmt.Println("The index for 3 is:", i)
		}
	}

	// range works with maps as well
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range and also return just the key or index
	for k := range kvs {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
