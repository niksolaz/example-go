package main

import "fmt"

type Fruit struct {
	Name  string
	Color string
}

type FruitMap map[string]Fruit

func exampleMapTwo(fruits []Fruit) FruitMap {
	mapFruit := make(FruitMap)
	for _, v := range fruits {
		mapFruit[v.Name] = v
	}
	return mapFruit
}

func updateMapFruiMap(fruitMap FruitMap, name string, color string) {
	fruitMap[name] = Fruit{name, color}
}

func deleteMapFruit(fruitMap FruitMap, name string) {
	delete(fruitMap, name)
}

func readMapFruit(fruitMap FruitMap, name string) Fruit {
	return fruitMap[name]
}

func addMapFruit(fruitMap FruitMap, name string, color string) {
	fruitMap[name] = Fruit{name, color}
}

func keyExist(fruitMap FruitMap, name string) bool {
	_, ok := fruitMap[name]
	return ok
}

func exampleMapOne() {
	// Vertex is a struct
	type Vertex struct {
		Lat, Long float64 // Latitude and Longitude of a location in float64
	}

	m := make(map[string]Vertex) // make a map of string to Vertex (struct) type

	m["Bell Labs"] = Vertex{ // key is "Bell Labs" and value is Vertex{40.68433, -74.39967}
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"]) // print the value of key "Bell Labs" from the map
	fmt.Println(m)              // print the map
}
func main() {
	exampleMapOne()

	arrFruits := []Fruit{
		{"Apple", "Red"},
		{"Banana", "Yellow"},
		{"Orange", "Orange"},
	}
	s := exampleMapTwo(arrFruits)
	fmt.Println(s)
	fmt.Println(s["Banana"].Color)

	updateMapFruiMap(s, "Banana", "Green")
	fmt.Println("update: ", s["Banana"])
	deleteMapFruit(s, "Banana")
	fmt.Println("delete: ", s)
	fmt.Println("read: ", readMapFruit(s, "Apple"))
	addMapFruit(s, "Ribes", "Blue")
	fmt.Println("add: ", s)
	fmt.Println("key exist: ", keyExist(s, "Ribes"))
}
