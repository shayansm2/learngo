# notes

### libraries:

1. you can compile a library package. but you cannot run it as it is not executable (does not have main package and main
   function)
2. if you want to have an exportable function or variable in a library which can be used by other packages, the first
   letter should be capital
3. in order to import your library, you should use the full path

```go
// Automatically imports!... AWESOME!
import "github.com/inancgumus/learngo/05-write-your-first-library-package/printer"
```