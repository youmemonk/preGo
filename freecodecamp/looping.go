package main

import "fmt"

func loopingFunc() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	fmt.Println("----------------")
	// s := []int{1, 2, 3}
	s := "Hello Go!"
	for k, v := range s {
		fmt.Println(k, string(v))
	}
}
