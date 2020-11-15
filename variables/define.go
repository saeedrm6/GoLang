package variables

import "fmt"

//How to Define variables?
func Defines() {
	// Deﬁne Variables

	// var variableName dataType = initialValu
	var integer1 int = 15
	var integer2 int8 = -25
	var integer3 int32 // default 0
	var float1 float32 = 63.2
	var string1 string = "Hello World!"
	var boolean1 bool // default false
	var boolean2 bool = true
	fmt.Println(integer1, integer2, integer3, float1, string1, boolean1, boolean2)

	// Type Inference
	// var variableName = initialValue
	var integer4 = 52           // int
	var string2 = "Hello World" // string
	var boolean3 = false        // bool
	fmt.Println(integer4, string2, boolean3)

	// Multiple Variable Declaration
	//var var1, var2, var3 int
	//var var1, var2, var3 int = 1, 2, 3
	//var var1, var2, var3 = 1, 2.2, false
	//var1, var2, var3 := 1, 2.2, false
	var5 := 15
	var (
		var1        = 50
		var2        = 25.22
		var3 string = "Telefonía"
		var4 bool
	)
	fmt.Println(var1, var2, var3, var4, var5)
}
