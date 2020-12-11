package interfaces

import "fmt"

func Defines() {
	/*
		Generic programing not supported in golang but some features like interface help us
		Define :    var i interface{}
		interface is kind of null type that can store types inside of her
		if inputs of function are interface, compilers sensitive be lower to check types and type errors
	*/

	var i interface{}
	describe1(i)

	i = 42
	describe1(i)

	i = "hello"
	describe1(i)

	i = []int{1, 2, 3}
	describe1(i)
}

func describe1(i interface{}) {
	fmt.Println("(%v, %T)\n", i, i)
}
