package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
    // Try using power of 0.5
    //x := 16
    //pwr := 0.5
    //fmt.Printf("Will this give square-root too? %v -> %g \n", x, x ** pwr)
    // Doesn't work - seems ** can't be used to raise to power, need math.Pow() instead
    
}

