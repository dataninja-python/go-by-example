package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 8}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum of 2, 3, 4, and 8:", sum)

	for ind, num := range nums {
		if num == 3 {
			fmt.Println("index of 3:", ind)
		}
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value)
	}

	for k := range kvs {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
