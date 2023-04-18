package package_doc

import (
	"fmt"
	"strings"
)

func StringDoc() {
	s := `hello
	how are you` // it's ` not '

	fmt.Println(strings.TrimSpace(s))

	char := "b"
	if strings.IndexAny(char, "aeiou") != -1 {
		fmt.Printf("Vowel")
	}

	fmt.Println(strings.ToUpper(strings.Repeat("yansha", 3)))
}
