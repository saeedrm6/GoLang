package main

import (
	"first/concurrency"
	"first/loopconditions"
	"first/oo_approach"
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

	// Loop & Conditions
	loopconditions.For()
	loopconditions.While()
	//loopconditions.Infinite()
	loopconditions.ForEachRange()
	loopconditions.ContinueBreak()
	loopconditions.IfElse()
	loopconditions.Switches()

	// Object Oriented vs Golang Approach
	oo_approach.Help()
	oo_approach.Methods()
	oo_approach.Embedding()
	oo_approach.Encapsulation()
	oo_approach.Example1()
	oo_approach.Example2()
	oo_approach.Example3()
	oo_approach.Pointer_And_Value_Receiver()

	// Concurrency
	concurrency.Goroutine_WaitGroup()
	concurrency.Goroutine_sleep()

}
