package main

import "fmt"

type person struct {
	name  string
	age   int
	email string
}

func addPerson(name string) *person {
	// This function returns a pointer to a person struct.
	fmt.Println("Adding person: ", name)
	fmt.Println("Unless more information is supplied, default values are used for age and email.")

	p := person{name: name}
	p.age = 0
	p.email = "none"
	return &p
}

func main() {
	fmt.Println("This file deals with structs.")
	fmt.Println("Structs are typed collections of fields.")
	fmt.Println("Grouped together they form records.")

	fmt.Println(person{"Bob", 20, "bob@gmail.com"})
	fmt.Println(person{name: "Alice", age: 30, email: "alice2023@me.com"})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(addPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
