package main

import "fmt"

// Struct - A collection of fields
type person struct { // struct - like a javascript object
	name string
	age  int
}

func main() {
	p := person{name: "Jake", age: 23}
	fmt.Println(p.age)
}
