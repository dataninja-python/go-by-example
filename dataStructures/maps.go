package main

import "fmt"

func main() {
	fmt.Println("Maps are Go's hashes or dicts")

	// Create a map
	m := make(map[string]int)
	fmt.Println("Create a map: m := make(map[string]int)")
	fmt.Println("This is a map with a string key and an int value.")
	fmt.Println("m:", m)

	// set key/value pairs
	fmt.Println("Set 'k1' <- 7")
	m["k1"] = 7
	fmt.Println("Set 'k2' <- 13")
	m["k2"] = 13
	fmt.Println("m:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	v3 := m["k3"]
	fmt.Println("v3: ", v3)

	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("m: ", m)

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)
}
