package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ name string }

func (cow Cow) Eat()   { fmt.Println("grass") }
func (cow Cow) Move()  { fmt.Println("walk") }
func (cow Cow) Speak() { fmt.Println("moo") }

type Bird struct{ name string }

func (bird Bird) Eat()   { fmt.Println("worms") }
func (bird Bird) Move()  { fmt.Println("fly") }
func (bird Bird) Speak() { fmt.Println("peep") }

type Snake struct{ name string }

func (snake Snake) Eat()   { fmt.Println("mice") }
func (snake Snake) Move()  { fmt.Println("slither") }
func (snake Snake) Speak() { fmt.Println("hsss") }

var animals map[string]Animal

type ExecutableCommand interface {
	execute()
}

type NewAnimalCommand struct {
	name       string
	animalType string
}

func (command NewAnimalCommand) execute() {
	var animal Animal
	switch command.animalType {
	case "cow":
		animal = Cow{name: command.name}
	case "bird":
		animal = Bird{name: command.name}
	case "snake":
		animal = Snake{name: command.name}
	default:
		panic("invalid animal type")
	}

	animals[command.name] = animal
	fmt.Println("Created it!")
}

type QueryCommand struct {
	name          string
	requestedInfo string
}

func (command QueryCommand) execute() {
	animal := animals[command.name]

	switch command.requestedInfo {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}

func main() {
	animals = make(map[string]Animal)
	for {
		var inputCommand, arg1, arg2 string
		fmt.Print(">")
		fmt.Scanln(&inputCommand, &arg1, &arg2)

		var command ExecutableCommand
		switch inputCommand {
		case "newanimal":
			command = NewAnimalCommand{name: arg1, animalType: arg2}
		case "query":
			command = QueryCommand{name: arg1, requestedInfo: arg2}
		}
		command.execute()
	}
}
