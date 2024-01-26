package main

import "fmt"

type FilterFunc func(int) bool
type MapperFunc func(int) int

func IsSquare(x int) bool {
	for i := 0; i*i <= x; i++ {
		if i*i == x {
			return true
		}
	}
	return false
}

func IsPalindrome(x int) bool {
	strForm := fmt.Sprintf("%d", Abs(x))
	for i := 0; i <= len(strForm)/2; i++ {
		if strForm[i] != strForm[len(strForm)-i-1] {
			return false
		}
	}
	return true
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Cube(num int) int {
	return num * num * num
}

func Filter(input []int, f FilterFunc) []int {
	result := make([]int, 0)
	for _, val := range input {
		if !f(val) {
			continue
		}
		result = append(result, val)
	}
	return result
}

func Map(input []int, m MapperFunc) []int {
	result := make([]int, len(input))
	for i, val := range input {
		result[i] = m(val)
	}
	return result
}
