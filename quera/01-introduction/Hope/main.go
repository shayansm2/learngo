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
	scanner.Scan()
	line := scanner.Text()

	args := strings.Split(line, " ")
	p, _ := strconv.Atoi(args[0])
	q, _ := strconv.Atoi(args[1])

	for i := 1; i <= q; i++ {
		if i%p != 0 {
			fmt.Println(i)
		} else {
			for j := 0; j < i/p; j++ {
				if j > 0 {
					fmt.Print(" ")
				}
				fmt.Print("Hope")
			}
			fmt.Println()
		}
	}
}
