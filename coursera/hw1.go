package main

import "fmt"

func main() {
	firstArray := [10]int{8, 2, 6, 1, 3, 9, 0, 5, 6, 3}
	secondArray := [10]int{8, -2, 6, 1, -3, 9, 0, 5, -6, 3}

	BubbleSort(firstArray[:])
	fmt.Println(firstArray) // [0 1 2 3 3 5 6 6 8 9]

	BubbleSort(secondArray[:])
	fmt.Println(secondArray) // [-6 -3 -2 0 1 3 5 6 8 9]
}

func BubbleSort(array []int) {
	for range array {
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				Swap(array, i)
			}
		}
	}
}

func Swap(array []int, index int) { array[index], array[index+1] = array[index+1], array[index] }
