package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (this Animal) Eat() {
	fmt.Println(this.food)
}

func (this Animal) Move() {
	fmt.Println(this.locomotion)
}

func (this Animal) Speak() {
	fmt.Println(this.noise)
}

func initAnimals() map[string]Animal {
	animals := make(map[string]Animal)
	animals["cow"] = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	animals["bird"] = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	animals["snake"] = Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	return animals
}

func main() {
	animals := initAnimals()

	for {
		var animalName, action string
		fmt.Print(">")
		fmt.Scanln(&animalName, &action)
		animal := animals[animalName]
		switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		}
	}
}
