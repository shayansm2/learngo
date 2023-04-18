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

	fmt.Println(strings.Repeat("yansha", 3))
	fmt.Println(strings.ToUpper("abc"))
	fmt.Println(strings.ToLower("DEF"))
	fmt.Println(strings.TrimSpace(" h i     j"))
	fmt.Println(strings.TrimRight("    k l m", " "))
}
