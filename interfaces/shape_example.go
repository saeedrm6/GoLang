package interfaces

import "fmt"

/*
	Shape agreement (signature method)
*/
type Shape interface {
	Area() int
}

/*
	Square Class and methods
*/
type Square struct {
	Length int
}

func (s Square) Area() int {
	return s.Length * s.Length
}

/*
	Rectangle Class and Methods
*/
type Rectangle struct {
	Length int
	Width  int
}

func (r Rectangle) Area() int {
	return r.Length * r.Width
}

func Example_1() {
	s := Square{Length: 10}
	r := Rectangle{Length: 10, Width: 30}

	shapeList := []Shape{s, r}

	for _, value := range shapeList {
		fmt.Printf("type : %T, Area : %d", value, value.Area())
	}
}
