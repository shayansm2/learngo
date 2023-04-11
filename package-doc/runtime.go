package package_doc

import (
	"fmt"
	"runtime"
)

func RunTimeDoc() {
	fmt.Println(
		runtime.NumCPU(),
		runtime.Version(),
	)
}
