package main

/*
Doing some things with pixel imd
*/

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	Width  = 1024
	Height = 768
)

func myshape(x, y float64, imd *imdraw.IMDraw) {
    // Draw some squares within squares on an exising imd
	imd.Color = pixel.RGB(1, 1, 0)
	for w := 0.0; w < 22; w = w + 2 {
		imd.Push(pixel.V(x-w, y+w))
		imd.Push(pixel.V(x+w, y+w))
		imd.Push(pixel.V(x+w, y-w))
		imd.Push(pixel.V(x-w, y-w))
		imd.Push(pixel.V(x-w, y+w))
		imd.Line(1)
	}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "things!",
		Bounds: pixel.R(0, 0, Width, Height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 0, 0)

    // Sort of Moirish pattern
	for x := 0.0; x <= Width; x = x + 10 {
		imd.Push(pixel.V(0, 0))
		imd.Push(pixel.V(x, Height))
		imd.Push(pixel.V(Width, 0))
		imd.Push(pixel.V(x, Height))
	}
	imd.Line(2)
	fmt.Println(reflect.TypeOf(imd))
	
    for !win.Closed() {
		win.Clear(pixel.RGB(1, 0, 1))
		imd.Draw(win)
        // Add some squares
		nx := rand.Float64() * Width
		ny := rand.Float64() * Height
		myshape(nx, ny, imd)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
