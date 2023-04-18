package package_doc

import (
	"fmt"
	"os"
	"strconv"
)

func strconvDoc() {
	f, _ := strconv.ParseFloat(os.Args[1], 64)
	fmt.Println(f)

	// Atoi: ASCII to integer
	number, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(number)

	// Itoa: Integer to ASCII
	ascii := strconv.Itoa(number)
	fmt.Println(ascii)

	fmt.Println("boolian status is " + strconv.FormatBool(true))
}
