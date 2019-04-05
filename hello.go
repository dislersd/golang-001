package main // if you want to execute any go file must have package main

import ( // import different packages here
	"errors"
	"fmt" // i/o package
	"math"
)

func main() { // func - go syntax - just like javascript - function name (params) {body}

	var x int = 5 // declare types -> int, bool, string, etc..
	fmt.Println(x)

	y := 10 // go can infer types - this is shorthand for storing a variable

	if y > 6 { // you can write if statement
		fmt.Println("More than 6")
	} else if y < 6 {
		fmt.Println("Less than 6")
	} else {
		fmt.Println("Broken")
	}

	a := []int{5, 4, 3, 2, 1} // setting an array
	a = append(a, 13)
	fmt.Println(a) // => {5, 4, 3, 2, 1, 13}

	vertices := make(map[string]int)
	// maps hold key value pairs like dictionaries in python
	// type def is 'map'['key types']'type of values' - use the built in 'make' function to create map
	/*
		'map'['keys']'values'

			map      string     int
		vertices["triangle"] = 2
	*/
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	//delete(vertices, "square")

	fmt.Println(vertices["triangle"]) // => 2

	for i := 0; i < 5; i++ { // standard for loop
		fmt.Println(i)
	}

	i := 0

	for i < 5 { // refactor for loop to become while loop
		fmt.Println(i)
		i++
	}

	arr := []string{"a", "b", "c"}

	for index, value := range arr { // using for on an array
		fmt.Println("index:", index, "value:", value)
		/*
			=>	index: 0 value: a
					index: 1 value: b
					index: 2 value: c
		*/
	}

	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m { // using for on a map
		fmt.Println("key:", key, "value:", value)
	}
	/*
		=>	key: a value: alpha
				key: b value: beta
	*/

	result := sum(2, 10) // invoking sum func defined below
	fmt.Println(result)

	result2, err := sqrt(13)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}
	p := person{name: "Jake", age: 23}
	fmt.Println(p)
}

func sum(x int, y int) int { // define new func here and invoke it within main func
	return x + y
}

func sqrt(x float64) (float64, error) { // finding square root of a number x - returns a defined error if given negative num
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}
