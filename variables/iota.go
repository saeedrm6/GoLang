package variables

import "fmt"

// set values to your variables when you want to increase them step by step from zero value

const (
	water = iota
	air
	soil
	glass
)

func Iota() {
	fmt.Println("\n-- IOTA --")
	fmt.Println(water, air, soil, glass)
}
