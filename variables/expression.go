package variables

import "fmt"

func Expression()  {
	var a = 11/2
	fmt.Println(a)  // what is 11? int => then return int type

	var b = 11.2/2
	fmt.Println(b)  // 11.2 is float32 => return float32
}