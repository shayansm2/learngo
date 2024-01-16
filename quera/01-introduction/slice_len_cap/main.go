package main

import "fmt"

func main() {
	slice := make([]int, 0)
	for i := 0; i < 20; i++ {
		fmt.Println(len(slice), cap(slice))
		slice = append(slice, i)
	}
	fmt.Println(len(slice), cap(slice))
}
