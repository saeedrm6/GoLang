package loopconditions

import "fmt"

func ContinueBreak() {
	sum := 0
	for i := 0; i < 5; i++ {
		if i%2 != 0 {
			continue
		}
		if i%2 != 1 {
			break
		}
		sum += i
	}
	fmt.Println(sum)

	// continue : redirect you to next loop without running next line codes
	// break : for stopping other loops to run
}
