package main

import "C"

import "math"


// Note the export line is important - also no space after //
// Function won't be added to the library without it!
//export add
func add(left, right int) int {
	return left + right
}


// Note the export line is important - also no space after //
//export sin
func sin(x float64) float64 {
	return math.Sin(x)
}

func main() {
}
