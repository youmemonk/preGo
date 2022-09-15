package main

import (
	"fmt"
	"math"
)

func controlStatementsFunc() {
	// if true {
	// 	fmt.Println("This test is true")
	// }

	number := 50
	guess := 64

	// if guess < 1 || returnTrue() || guess > 100 { //short circuiting
	if guess < 1 || guess > 100 { //short circuiting
		fmt.Println("The guess must be in between 1 and 100")
	} else {
		if guess < number {
			fmt.Println("Too Low")
		}
		if guess > number {
			fmt.Println("Too High")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
		fmt.Println(guess >= number, guess <= number, guess != number)
	}

	// float32Comparison()
	// switchComparision()
	switchTypeComparison()

}

func returnTrue() bool {
	fmt.Println("Returning True")
	return true
}

func float32Comparison() {
	myNum := 0.123
	// if myNum == math.Pow(math.Sqrt(myNum), 2) {
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are Same")
	} else {
		fmt.Println("These are different")
	}
}

func switchComparision() {
	// switch i := 2 + 3; i {
	i := 10
	switch {
	// case 1, 5, 10:
	// 	fmt.Println("one, five or ten")
	// case 2, 4, 6:
	// 	fmt.Println("two, four or six")
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough
	case i >= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("Another number")
	}
}

func switchTypeComparison() {
	// var i interface{} = [2]int{1, 2}
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		break
		fmt.Println("This is after the break")
	case float64:
		fmt.Println("i is float64")
	case string:
		fmt.Println("i is string")
	case [2]int:
		fmt.Println("i is []int slice")
	default:
		fmt.Println("i is another type")
	}

}
