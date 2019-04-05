package main // if you want to execute any go file must have package main

import ( // import different packages here
	"fmt" // i/o package
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

}
