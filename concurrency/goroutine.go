package concurrency

import (
	"fmt"
	"sync"
	"time"
)

/*
	GoRoutines :
	go function name
	your program will be run with two branches
	* in goroutines you always need a waiter like time.sleep() or sync.WaitGroup
*/

func say_syncWaitGroup(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		//time.Sleep(1 * time.Millisecond)
		fmt.Println(s)
	}
	wg.Done()
}

func say_sleep(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(s)
	}
}

func Goroutine_WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	go say_syncWaitGroup("world", &wg)
	wg.Wait()
	wg.Add(1)
	say_syncWaitGroup("hello", &wg)
	wg.Wait()
	fmt.Println("END of Go Routine WaitGroup")
}

func Goroutine_sleep() {
	go say_sleep("world")
	say_sleep("hello")
	fmt.Println("END of Go Routine Sleep")
}
