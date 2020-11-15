package variables

import (
	"fmt"
	"math"
)

//Type Conversion
func Converstion() {
	var1 := 10   // int
	var2 := 10.5 // float64
	// illegal
	//var3 := var1 + var2
	//legal
	var3 := var1 + int(var2) // var3 => 20
	fmt.Println(var3)

	var (
		var4 float64 = -1.2
		var5 int     = 1
	)
	var6 := var4 + float64(var5)
	if var6 == -0.2 {
		fmt.Println("It is Equal")
	}
	const float64EqualityThresold = 1e-9
	if math.Abs(var6-(-0.2)) < float64EqualityThresold {
		fmt.Println("It is Equal")
	}
	fmt.Println(var6)
}

// calculating floats are never secure in if conditions
// IEEE754
