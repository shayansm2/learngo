package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		name := scanner.Text()

		scanner.Scan()
		line := scanner.Text()
		scores := strings.Split(line, " ")
		var sum, count int
		for _, score := range scores {
			intScore, _ := strconv.Atoi(score)
			sum += intScore
			count++
		}
		avg := sum / count
		var label string
		switch {
		case avg >= 80:
			label = "Excellent"
		case avg >= 60:
			label = "Very Good"
		case avg >= 40:
			label = "Good"
		default:
			label = "Fair"
		}

		fmt.Printf("%v %v\n", name, label)
	}
}
