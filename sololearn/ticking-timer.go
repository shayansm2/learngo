package main

import "fmt"

type Timer struct {
	id    string
	value int
}

func (p *Timer) tick() {
	p.value++
	fmt.Println(p.value)
}

func main() {
	var x int
	fmt.Scanln(&x)

	t := Timer{"timer1", 0}

	for i := 0; i < x; i++ {
		t.tick()
	}
}
