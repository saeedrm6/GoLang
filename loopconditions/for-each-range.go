package loopconditions

import "fmt"

func ForEachRange() {
	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
}
