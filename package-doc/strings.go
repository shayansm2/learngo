package package_doc

import (
	"fmt"
	"strings"
)

func StringDoc() {
	var s string
	s = `
	hello
	how are you
	`
	fmt.Println(strings.TrimSpace(s))

	char := "b"
	if strings.IndexAny(char, "aeiou") != -1 {
		fmt.Printf("Vowel")
	}

	fmt.Println(strings.ToUpper(strings.Repeat("yansha", 3)))
}
