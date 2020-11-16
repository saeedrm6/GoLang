package structures

import "fmt"

func Arrays() {
	// Array: Store fixed number of elements
	// var myArray [size]type
	// var integerArray [5]int

	a := [5]int{1, 2, 3, 4, 5}
	b := [2]string{"first", "second"}

	fmt.Println(a, b)

}
