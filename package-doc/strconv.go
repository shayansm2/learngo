package package_doc

import (
	"fmt"
	"os"
	"strconv"
)

func strconvDoc() {
	f, _ := strconv.ParseFloat(os.Args[1], 64)
	fmt.Println(f)
}
