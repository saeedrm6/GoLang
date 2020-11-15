package variables

import "fmt"

func Pointers() {
	i := 52
	var p *int
	p = &i

	// & give u the memory address and * give you the corrent value of target variable
	fmt.Println(i, p, &i, *p)

	i = 102
	fmt.Println(i, *p)
}
