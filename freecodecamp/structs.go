package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
	episodes   []string
}

type Animal struct {
	Name   string `required max: "100"` //tag
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func AnimalStruct() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48.50
	b.CanFly = false

	anotherBird := Bird{
		Animal:   Animal{Name: "Ostrich", Origin: "Australia"},
		SpeedKPH: 24,
		CanFly:   false,
	}

	fmt.Println(b)
	fmt.Println(b.Name)

	fmt.Println(anotherBird)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

func structsFunc() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Clark Peterson",
		companions: []string{
			"Jo Zhee",
			"Jo Grant",
			"Shawn Mendes",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[2])

	anotherDoctor := aDoctor
	anotherDoctor.actorName = "Miachel B. Jordan"

	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)

	AnimalStruct()
}
