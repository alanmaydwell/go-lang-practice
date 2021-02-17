package main

/*
Draw some moving things
*/

import (
	"fmt"

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

// Change to have a slice of points to plot relative to x1, y1
type Movable struct {
	x1, y1 float64
	x2, y2 float64
	dx, dy float64
}

// Pointer needed here for updates to work
func (m *Movable) move(){
    m.x1 += m.dx
    m.y1 += m.dy
    m.x2 += m.dx
    m.y2 += m.dy
}


func run() {
	// Create pixelgl Window
	cfg := pixelgl.WindowConfig{
		Title:  "Hmm?",
		Bounds: pixel.R(0, 0, Width, Height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	var things []Movable
	for x := 0.0; x < 16.0; x++ {
		thing := Movable{x1: 0, y1: 30 * x, x2: 20, y2: 32 * x, dx: 1, dy: 0.3}
		things = append(things, thing)
	}

	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(0, 1, 1)

	fmt.Println("things")

	for !win.Closed() {

		// Screen clear and imd clear, need both
		win.Clear(pixel.RGB(0, 0, 0))
		imd.Clear()

		// Draw things
		for ti, thing := range things {
			fmt.Println(thing.x1, thing.y1)
			imd.Push(pixel.V(thing.x1, thing.y1))
			imd.Push(pixel.V(thing.x2, thing.y2))
			imd.Line(1)
            things[ti].move()
		}

		imd.Draw(win)
		// Update the window to make changes visible
		win.Update()

	}
}

func main() {
	pixelgl.Run(run)
}
