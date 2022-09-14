package main

import (
	"fmt"
)

// const a int16 = 12
const (
	// iota is scoped to this constant block
	a = iota
	b
	c
)

const (
	//here the iota resets it's value, as its scoped to a different constant block now
	a2 = iota
)

const (
	catSpecialist = iota
	dogSpecialist
	snakeSpecialist
)

func constants() {
	var specialistType int = dogSpecialist
	fmt.Printf("%v\n", specialistType == dogSpecialist)

	// fmt.Printf("%v, %T\n", a, a)
	// fmt.Printf("%v, %T\n", b, b)
	// fmt.Printf("%v, %T\n", c, c)

	// fmt.Printf("%v, %T\n", a2, a2)

	// const myConst int = 42
	// myConst = 27 // not possible to change the type

	// const a = 29 //shadowing
	// var b int16 = 12
	// fmt.Printf("%v %T\n", a, a)

	// fmt.Printf("%v %T\n", a+b, a+b)
	//test
}
