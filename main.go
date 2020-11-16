package main

import (
	"first/printer"
	"first/structures"
	"first/variables"
)

func main() {
	// Add and user packages
	printer.PrintHelloWorld()

	// Variables
	variables.Defines()
	variables.Converstion()
	variables.Aliasing()
	variables.Iota()
	variables.Expression()
	variables.Pointers()

	// Data Structures
	structures.Maps()
	structures.Arrays()
	structures.Slices()
	structures.Types()
}
