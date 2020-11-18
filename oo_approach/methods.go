package oo_approach

import "fmt"

func (c Creature) Dump() {
	fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}

func Methods() {
	/*
		func (c Creature) Dump() {
			fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
		}
	*/

	var testObj Creature
	testObj.Name = "Saeed"
	testObj.Real = true
	fmt.Println(testObj)
	testObj.Dump()
}
