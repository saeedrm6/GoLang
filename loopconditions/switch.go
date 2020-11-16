package loopconditions

import "fmt"

func Switches() {
	dayOfWeek := 2
	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
		break
	case 2:
		fmt.Println("two")
		fmt.Println("another two")
		break
	default:
		fmt.Println("no matches")

	}
}
