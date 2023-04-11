package package_doc

import (
	"fmt"
	"path"
)

func PathDoc() {
	dir, file := path.Split("css/main.css")
	fmt.Println(dir, file)
}
