package printer

import (
	"fmt"
	"time"
)

//Print Message of intoruction
func PrintHelloWorld() {
	beginTime()
	fmt.Println("Hello World!")
}

//Print Current DATETIME
func beginTime() {
	fmt.Println("DATETIME: " + time.Now().Format("2006-01-02 15:04:05 MST"))
}
