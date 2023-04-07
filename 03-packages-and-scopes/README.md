# notes

### packages

- You can access functions from other files which are in the same package

### scopes

we have three scopes in go which are:

1. package scope
2. file scope
3. block scope

- everything you define in a function is in **block scope**
- importing packages in a file is in **file scope**
- defining variables and functions in a package is in **package scope**
- you cannot declare a function or a variable with the same name and the same scope
- you can do this with same name but different scopes. The inner scope definition will override in the inner scope

### renaming imports

```go
import f "fmt"
```