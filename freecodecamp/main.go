package main

import (
	"fmt"
	"strconv"
)

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

func main() {
	var i int
	i = 97
	fmt.Printf("%v %T", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Println(j)
	fmt.Printf("%v %T", j, j)
}
