package main

import "fmt"

func mapsFunc() {
	statePopulation := make(map[string]int, 10)
	statePopulation = map[string]int{
		"California": 234500,
		"Texas":      123445,
		"Nevada":     982459,
		"Arizona":    234921,
	}

	fmt.Println(statePopulation)
	fmt.Println(statePopulation["Texas"])

	statePopulation["Ohio"] = 696969
	statePopulation["Pennisilvenia"] = 235456

	// fmt.Println(statePopulation)
	// delete(statePopulation, "Texas")
	// fmt.Println(statePopulation)

	// pop, ok := statePopulation["California"]
	_, ok := statePopulation["California"]
	fmt.Println(ok)
	// fmt.Println(pop, ok)
	// m := map[[2]int]string{}

	// fmt.Println(m)

	fmt.Println(len(statePopulation))
	fmt.Println(statePopulation)

	sp := statePopulation
	delete(statePopulation, "Texas")
	fmt.Println(statePopulation)
	fmt.Println(sp)

	if pop, ok := statePopulation["Arizona"]; ok {
		fmt.Println(pop)
	}
}
