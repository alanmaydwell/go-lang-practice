package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("The fates have decreed it's a:", rand.Intn(10))
    fmt.Println("or maybe a:", rand.Intn(10))
    x := rand.Intn(10)
    fmt.Printf("unless it's %v instead (or possibly as well)!\n", x)
}

