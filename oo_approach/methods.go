package oo_approach

import "fmt"

/*
	Dump belongs to Creature Struct Type ==> before defining name of function (variableName TypeName)
*/
func (c Creature) Dump() {
	fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}

/*
	Simple function with input string
*/
func sayHi(name string) {
	fmt.Println(name)
}

/*
	Simple function with input string and determine return type
*/
func sayHello(name string) bool {
	return true
}

func Methods() {
	var testObj Creature
	testObj.Name = "Saeed"
	testObj.Real = true
	fmt.Println(testObj)
	testObj.Dump()

	sayHi(testObj.Name)
	sayHello("saeed")
}
