package main

import (
	"fmt"
)

type DiscreteAnimal struct {
	food       string
	locomotion string
	noise      string
}

func (this DiscreteAnimal) Eat() {
	fmt.Println(this.food)
}

func (this DiscreteAnimal) Move() {
	fmt.Println(this.locomotion)
}

func (this DiscreteAnimal) Speak() {
	fmt.Println(this.noise)
}

func initAnimals() map[string]DiscreteAnimal {
	animals := make(map[string]DiscreteAnimal)
	animals["cow"] = DiscreteAnimal{food: "grass", locomotion: "walk", noise: "moo"}
	animals["bird"] = DiscreteAnimal{food: "worms", locomotion: "fly", noise: "peep"}
	animals["snake"] = DiscreteAnimal{food: "mice", locomotion: "slither", noise: "hsss"}
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
