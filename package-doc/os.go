package package_doc

import (
	"fmt"
	"os"
)

func OsDoc() {
	fmt.Println(os.Args, os.Args[1])
}
