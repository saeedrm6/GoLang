package variables

import "fmt"

// متغیر هایی که میخوایی یه دونه یه دونه مقدار بگیرند

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
