# notes

- use `package main` To create an executable Go program
- using `func main` allows Go to start executing the program
- you do not call the main function by yourself. Go calls the main function automatically.
---
`go run` both compiles and runs a program; whereas `go build` just compiles it
1. Run a Go program: `go run main.go`
2. - Build a Go program: `go build main.go`
   - and then run it: `./main`


If you have multiple files in a package: `go run .` or `go run *.go`