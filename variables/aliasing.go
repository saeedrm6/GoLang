package variables

import "fmt"

type floatsaeed float64

func Aliasing() {
	// type aliasName aliasTo
	var f floatsaeed = 1.24456
	fmt.Printf("f has value %v and type %T", f, f)

	// Note : you cant compare different variables with non equeal data type even your new type is same of other variable data type
}
