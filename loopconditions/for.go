package loopconditions

import "fmt"

func For() {
	/*
				-->condition
		for --  -->for caluse (init; cond; post)   --> { Block }
				-->range

	*/

	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println("Sum is ", sum)

	// You cant use variable i outside of for
}
