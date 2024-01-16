package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n int
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	countryCodes := make(map[string]string)

	for i := 0; i < n; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		countryCodes[input[1]] = input[0]
	}

	var q int
	scanner.Scan()
	q, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < q; i++ {
		scanner.Scan()
		country, found := countryCodes[scanner.Text()[:3]]
		if found {
			fmt.Println(country)
		} else {
			fmt.Println("Invalid Number")
		}
	}
}
