package main

import (
	"fmt"
	"time"
)

func main() {
	var ball int
	table := make(chan int)
	go player(table, "player1")
	go player(table, "player2")
	//go player(table, "player3")

	table <- ball
	time.Sleep(1 * time.Second)
	<-table
}

func player(table chan int, playerName string) {
	for {
		ball := <-table
		ball++
		fmt.Println(playerName, "hit the", ball)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
