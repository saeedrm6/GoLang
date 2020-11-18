package oo_approach

import "fmt"

type foo struct {
}

func (f foo) Foo1() {
	fmt.Println("Foo1() here")
}

func (f foo) Foo2() {
	fmt.Println("Foo2() here")
}

func (f foo) Foo3() {
	fmt.Println("Foo3() here")
}

var Footest = foo{}

func Encapsulation()  {
	Footest.Foo1()
	Footest.Foo2()
	Footest.Foo3()
}