package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	library := make(map[int]string)
	for i := 0; i < n; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		command := input[0]
		isbn, _ := strconv.Atoi(input[1])
		switch command {
		case "ADD":
			bookName := strings.Join(input[2:], " ")
			library[isbn] = bookName
		case "REMOVE":
			delete(library, isbn)
		}
	}

	isbns := make([]int, len(library))
	i := 0
	for isbn := range library {
		isbns[i] = isbn
		i++
	}

	sort.Slice(isbns, func(i, j int) bool {
		if library[isbns[i]] == library[isbns[j]] {
			return isbns[i] < isbns[j]
		}
		return library[isbns[i]] < library[isbns[j]]
	})

	for _, isbn := range isbns {
		fmt.Println(isbn)
	}
}
