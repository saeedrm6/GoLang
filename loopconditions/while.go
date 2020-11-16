package loopconditions

import "fmt"

func While() {
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println(n)
}
