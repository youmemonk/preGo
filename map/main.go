package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b3945",
		"white": "#ffffff",
	}

	// var colors map[string]string

	// colors := make(map[int]string)

	// colors[10] = "#ffffff"

	// delete(colors, 10)

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for the color", color, "is", hex)
	}
}
