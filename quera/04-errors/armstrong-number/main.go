package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

func main() {
	var input string
	fmt.Scanln(&input)
	if isArmstrong(getNumberFromStr(input)) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func getNumberFromStr(s string) int {
	var result, startIndex int
	for i, char := range s {
		if unicode.IsDigit(char) {
			continue
		}
		if i-1 >= startIndex {
			num, _ := strconv.Atoi(s[startIndex:i])
			result += num
		}
		startIndex = i + 1
	}

	if len(s) > startIndex {
		num, _ := strconv.Atoi(s[startIndex:])
		result += num
	}

	return result
}

func isArmstrong(num int) bool {
	length := len(strconv.Itoa(num))
	sum := 0
	for _, char := range strconv.Itoa(num) {
		digit, _ := strconv.Atoi(string(char))
		sum += int(math.Pow(float64(digit), float64(length)))
	}
	return sum == num
}
