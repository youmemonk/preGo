package main

import (
	"fmt"
)

func arrayfunc() {
	// a := []int{1, 2, 3}
	// b := a

	// b[1] = 5

	// fmt.Println(a)
	// fmt.Println(b)

	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v", cap(a))

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%v\n", a)

	// b := a[:]
	// c := a[3:]
	// d := a[:3]
	// e := a[2:8]

	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	fmt.Printf("Length %v\n", len(a))
	fmt.Printf("Capacity %v\n", cap(a))

	a = append(a, 1)
	fmt.Println(a)

	a = append(a, []int{11, 12, 13}...)
	fmt.Println(a)

	b := a[1:]
	c := a[:len(a)-1]
	fmt.Println(b)
	fmt.Println(c)

	d := append(a[:2], a[4:]...)
	fmt.Printf("Slice d : %v\n", d)
}
