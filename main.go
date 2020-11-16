package main

import (
	"first/printer"
	"first/structures"
	"first/variables"
	"first/loopconditions"
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

	// Loop & Conditions
	loopconditions.For()
	loopconditions.While()
}
