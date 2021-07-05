package interfaces

import "fmt"

/*
Type Assertion
	A type assertion provides access to an interface value's underlying concrete value.
	t, ok := i.(T)
*/

func Example_2() {
	s := Square{Length: 50}
	r := Rectangle{Length: 50, Width: 30}

	shapeList := []Shape{s, r}

	for _, value := range shapeList {
		fmt.Printf("\ntype : %T, Area : %d \n", value, value.Area())

		square, ok := value.(Square)
		fmt.Printf("type: %T, OK: %v\n", square, ok)
	}
}
