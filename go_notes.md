# Go Notes
Misc notes/cheat sheet on using Go

Note handy and interactive intro to Go here:
https://tour.golang.org/basics/1

## Minimal Start

```
package main

import "fmt"

func main() {
	fmt.Println("Greatings earthling!")
}
```

Note
   
- `package main` declaration at top   
- imported packages need quotation marks ("fmt"), unlike Python module import.
- Need `func main()`

## Compilation from terminal

#### Compile to binary
`go build <source filename>`

#### Compile and immediately execute 
`go run <source filename>`

## Declarations
Within functions it's not necessary to declare variables - can use shortcut 
using := operator (walrus?). Works a bit like duck typing in Python. 

```
func mything() {
	x := 1
	yf := 2.0
	name := "Galufrax"
    z := float64(x) * yf
    fmt.Println(name, z)
}
```
*Aside, needed to use float64(*) function to calculate z because
*can't multiply int and float.*

Note this is **not** allowed outside function
definitions. var statement can be used to declare variables inside or outside functions.

```
var yesno bool
// If consecutive variable share the same type, only need to declare it at the end
var flag1, flag2, flag3 bool
// Can also assign values at the same time as declaring
var x, y int = 1, 2
// If values are assigned it's not necessary to declare the type which can be taken from initial value instead
var a, b = 1, 2
``` 

### What is the type of a thing !!???
When defining a function's arguments and return value(s) it's necessary
to specify the type of each item, this is obvious in some cases such as int,
but what about more exotic things that have been imported from a library package?    

Can find the type name using reflect.Typeof() function.

First need to import reflect package

`import reflect`

Then later

`fmt.Println(reflect.TypeOf(strangething))`




## For loops (the only type of loop available)
Superficially like C but without need for brackets around the loop components.

```
for i := 0; i <= 10; i++ {
	fmt.Prinln(i)
}
```

Also has some extra versatility as can be used like a while loop (could have retained the semicolons here but seems little point).

```
total := 0
for total < 100 {
	total += 1
}

```

Can also loop forever as follows:

```
for {

}
```

