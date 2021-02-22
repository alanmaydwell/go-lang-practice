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
