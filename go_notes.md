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
## Else
**if**, **else if**, and **else** seem fairly straightforward. Couple of things to note:

- no brackets () around conditions
- **else** statement must be on the same line as its associated if's closing brace
i.e. `} else {`

## Map
Maps are similar to dictionaries. Can define like this:

```
things := map[string][]int{
    "This":{1,2,3},
    "That":{2,4,6},
    "Other":{2,4,8}
}
```
[string] above defines the data type of the keys and []int the data type of the values.

Can iterate over contents using range

`for key, value := range things {`

## A bit about structs and methods
```
package main

// Example of a struct with a method that updates its fields
// and another method that displays two of them

import "fmt"

type Movable struct {
	x, y int
	dx, dy int
    name string
}

// Method created like this:
//func (receiver receiver_type) some_func_name(arguments) return_values
// receiver is a bit like self in Python

// Note using pointer *Movable rather than Movable otherwise
// the updates below don't actually affect the supplied instance
func (m *Movable) move(){
    m.x += m.dx
    m.y += m.dy
}

// This method just displays details without update so pointer not needed here
// although could still be used but as we're using temporary substitution here
// for empty name, we don't want to update the original
func (m Movable) show_coords(){
    // Swap in a name if name is blank
    if m.name == "" {
        m.name = "NO NAME!"
    }
    fmt.Printf("%v - x: %v, y: %v\n", m.name, m.x, m.y)
}

func main() {
    // can create instance like this
    m1 := Movable{10, 10, 1, 0, "Argle"}
    // or like this
    m2 := Movable{x: 20, y:20, dx: 0, dy: 1, name: "Zoot"}
    // Don't need to specify each field - defaults to zero values
    m3 := Movable{}

    for {
        m1.show_coords()
        m2.show_coords()
        m3.show_coords()
        m1.move()
        m2.move()
        m3.move()
    }

}
```