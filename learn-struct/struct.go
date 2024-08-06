package main

import (
	"fmt"
	"image/color"
)

type Person struct {
	name string
	age  int
}

func GetPerson() {
	person := Person{"John", 25}
	fmt.Println(person)
	fmt.Println(person.name)
	fmt.Println(person.age)
}

type Fruit struct {
	name   string
	color  color.RGBA
	season string
}

func GetFruit(fruit Fruit) {
	fmt.Println(fruit)        // {Apple {255 0 0 255} Winter}
	fmt.Println(fruit.name)   // Apple
	fmt.Println(fruit.color)  // {255 0 0 255}
	fmt.Println(fruit.season) // Winter
}

func main() {
	GetPerson()
	GetFruit(Fruit{"Apple", color.RGBA{255, 0, 0, 255}, "Winter"})

	v := Fruit{"Banana", color.RGBA{255, 255, 0, 255}, "Winter"}
	p := &v
	p.name = "Mango"
	fmt.Println("Result of pointers struct", v)
}
