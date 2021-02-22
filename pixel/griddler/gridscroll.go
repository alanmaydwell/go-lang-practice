package main

/*
Used pixel imdraw to create a grid of lines accross whole window.
Gird can scroll in 8 directions
*/

import (
	"fmt"
	"reflect"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	// Window dimensions
	Width   = 1024
	Height  = 768
	Fwidth  = float64(Width)
	Fheight = float64(Height)
)

func liner(xspacing, xoffset int, vertical bool, imd *imdraw.IMDraw) {
	// Create vertical or horizontal lines at given spacing accross whole window
	imd.Color = pixel.RGB(0, 0, 1)
	for x := xoffset; x <= Width; x += xspacing {
		fx := float64(x) // need float as pixel.V only uses floats
		if vertical {
			imd.Push(pixel.V(fx, 0.0))
			imd.Push(pixel.V(fx, Fheight))
			// Go trap - the else needs to be on same line as the if's closing }
		} else {
			imd.Push(pixel.V(0.0, fx))
			imd.Push(pixel.V(Fwidth, fx))
		}
		imd.Line(1)
	}

}

func run() {
	// Create pixelgl Window
	cfg := pixelgl.WindowConfig{
		Title:  "Lines, Lines, Lines!",
		Bounds: pixel.R(0, 0, Width, Height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	fmt.Println(reflect.TypeOf(imd))

	// Width of grid squares
	spacing := 32

	directions := [][2]int{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1}}

	counter := 0
	xshift := 0
	yshift := 0
	direction := 7

	for !win.Closed() {

		if counter%128 == 0 {
			direction = (direction + 1) % len(directions)
			fmt.Printf("counter: %v, imode: %v\n", counter, direction)
		}

		// Screen update/animate (1) Wipe
		win.Clear(pixel.RGB(0, 0, 0))
		// Need to clear the imd as well as the window
		imd.Clear()
		// Redraw the grid
		// vertical lines
		liner(spacing, xshift, true, imd)
		// horizontal lines
		liner(spacing, yshift, false, imd)
		imd.Draw(win)
		// Update the window to make changes visible
		win.Update()

		// Scroll in current direction for a while
		xshift = (xshift + directions[direction][0]) % spacing
		yshift = (yshift + directions[direction][1]) % spacing
		counter++
	}
}

func main() {
	pixelgl.Run(run)
}
