# notes:

## data types:

**int , float32 , float64 , bool , string , etc ...**

```go
// integer types
var i int
var i8 int8
var i16 int16
var i32 int32
var i64 int64

// float types
var f32 float32
var f64 float64

// complex types
var c64 complex64
var c128 complex128

// bool type
var b bool

// string types
var s string
var r rune // also a numeric type
var by byte // also a numeric type
```

> Go is a statically-typed / Strongly-Typed language

## variables

### declaration

```go
var speed int    // numeric type
var heat float64 // numeric type
var off bool
var brand string

fmt.Println(speed) // 0
fmt.Println(heat)  // 0
fmt.Println(off) // false

// used printf to print an empty string
fmt.Printf("%q\n", brand) // ""
```

- you can not declare a variable and do not use it in block scope (gets an error)
- however you can do it in package scope (only gets a warning)
- you can use blank identifier in order to not get an error in block scope

```go
var speed int
_ = speed
```

### declaration with value

```go
// option 1 (so so)
var safe bool = true
// -------------------------------------------------------
// option 2 (ok) For package scoped variable
// also suitable when you don't know the initial value
var safe = true
// -------------------------------------------------------
// option 3 (best) Can't be used in package scope
// also suitable when you do know the initial value 
safe := true
```

### multiple declaration

```go
var (
speed int
heat  float64
off   bool
brand string
)
// ----------------------------------
var speed, velocity int
// ----------------------------------
safe, speed := true, 50

```

### value assignment
just like other languages
```go
var speed int
speed = 100
speed = speed - 25
// ----------------------------------
speed, now = 100, time.Now()
// ----------------------------------
speed, prevSpeed = prevSpeed, speed
```