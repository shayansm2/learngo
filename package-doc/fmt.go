package package_doc

import "fmt"

func FmtDoc() {
	fmt.Println("anything")
	fmt.Printf("%q\n", "")

	var input string
	fmt.Scanln(&input)
}
