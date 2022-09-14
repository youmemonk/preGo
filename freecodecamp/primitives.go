package main

import "fmt"

var i int = 32

var (
	actorName     string = "Elizabeth Olsen"
	companionName string = "Sarah Jane Smith"
	doctorNumber  int    = 3
	seasonNo      int    = 4
)

var (
	counter int = 25
)

func primitives() {

	r := 'a'
	fmt.Printf("%v %T\n", r, r)

	// s := "this is a string"
	// b := []byte(s)

	// fmt.Printf("%v %T\n", s, s)
	// fmt.Printf("%v %T\n", b, b)

	// a := 10 //1010
	// b := 3  //0011

	// fmt.Println(a & b)  // 0010 -> 2
	// fmt.Println(a | b)  // 1011 -> 11
	// fmt.Println(a ^ b)  // 1001 -> 9
	// fmt.Println(a &^ b) // 0100 -> 4

	// var i int
	// i = 97
	// fmt.Printf("%v %T", i, i)

	// var j string
	// j = strconv.Itoa(i)
	// fmt.Println(j)
	// fmt.Printf("%v %T", j, j)
}
