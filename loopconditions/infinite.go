package loopconditions

import "fmt"

func Infinite() {
	sum := 0
	for {
		sum++
	}
	fmt.Println(sum)
}
