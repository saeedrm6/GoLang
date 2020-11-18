package oo_approach

type FlyingCreature struct {
	Creature
	WingSpan int
}

/*
	"," end of every field in defining is important
*/

func Embedding() {
	dragon := &FlyingCreature{
		Creature{"Dragon", false},
		15,
	}
	dragon.Dump()
}
