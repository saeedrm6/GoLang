package concurrency

import (
	"fmt"
	"sync"
)

/*
	Channels :
    Channels are a typed conduit through which you can
	send and receive values with the channel operator

	ch := make(chan int)
	ch <- v  send v to channel ch
	v := <-ch receive from ch and assign value to v
*/

func sayHello(wg *sync.WaitGroup,pingChannel chan string,pongChannel chan string) {
	wg.Done()
	for i:=0;i<5;i++ {
		fmt.Println("Hello ")
		pingChannel <- "World!"
		<-pongChannel
	}
}

func sayWorld(wg *sync.WaitGroup,pingChannel chan string,pongChannel chan string) {
	wg.Done()
	for i:=0;i<5;i++ {
		v := <-pingChannel
		fmt.Println(v)
		pongChannel <- "Pong"
	}
}

func Channels()  {
	var wg sync.WaitGroup
	pingChannel := make(chan string)
	pongChannel := make(chan string)

	wg.Add(1)
	go sayHello(&wg,pingChannel,pongChannel)
	wg.Wait()

	wg.Add(1)
	go sayWorld(&wg,pingChannel,pongChannel)
	wg.Wait()

	fmt.Println("Finished")
}
