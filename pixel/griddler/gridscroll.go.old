package main

/*
Used pixel imdraw to create a grid of lines accross whole window.
Gird can scroll in 8 directions

Note coordinates are float64 because uses pixel.V to do draw to screen and it
requires float64
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

func liner(xspacing, xoffset float64, vertical bool, imd *imdraw.IMDraw) {
	// Create vertical or horizontal lines at given spacing accross whole window
	imd.Color = pixel.RGB(0, 0, 1)
	for x := xoffset; x <= Width; x += xspacing {
		if vertical {
			imd.Push(pixel.V(x, 0.0))
			imd.Push(pixel.V(x, Fheight))
			// Go trap - the else needs to be on same line as the if's closing }
		} else {
			imd.Push(pixel.V(0.0, x))
			imd.Push(pixel.V(Fwidth, x))
		}
		imd.Line(1)
	}

}

func xyshifts(spacing, shift float64, mode int) (xshift, yshift float64) {
	// Return scroll direction x/y shift parameters for given direction mode
	switch mode {
	// Stationary
	case 0:
		xshift = 0
		yshift = 0
	// Move right
	case 1:
		xshift = shift
		yshift = 0
	// Move left
	case 2:
		xshift = spacing - shift
		yshift = 0
	// Move up
	case 3:
		xshift = 0
		yshift = shift
	// Move down
	case 4:
		xshift = 0
		yshift = spacing - shift
	// diag NE
	case 5:
		xshift = shift
		yshift = shift
	// diag NW
	case 6:
		xshift = spacing - shift
		yshift = shift
	// diag SE
	case 7:
		xshift = shift
		yshift = spacing - shift
	case 8:
		// diag SW
		xshift = spacing - shift
		yshift = spacing - shift
	default:
		xshift = 0
		yshift = 0
	}
	return xshift, yshift
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
	spacing := 32.0
	direction_sequence := []int{1, 7, 4, 8, 2, 6, 3, 5}

	counter := 0
	imode := 0
	mode := direction_sequence[0]
	for !win.Closed() {

		if counter%8 == 0 {
			imode = (imode + 1) % len(direction_sequence)
			mode = direction_sequence[imode]
			fmt.Printf("counter: %v, imode: %v, mode: %v\n", counter, imode, mode)
		}

		// Scroll smoothy accross width/height of grid cell
		for shift := 0.0; shift < spacing; shift++ {
			xshift, yshift := xyshifts(spacing, shift, mode)
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
		}
		counter++
	}

}

func main() {
	pixelgl.Run(run)
}
