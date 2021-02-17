package main

/*
Draw some shapes based on points on edge of circle
*/

import (
	"fmt"
	"math"

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

/*
Make and return array of x/y points that sit on the circumference of a circle
  	angle1 - start angle (radians)
 	angle2 - end angle (radians)
  	dangle - delta angle (radians)
   radius - what is says (pixels but float)
Coordinates are float64 although intended for drawing as pixels as want
to draw using pixel.V (vector) which requires float64.
*/
func circle_points(angle1, angle2, dangle, radius float64) [][2]float64 {
	var coords [][2]float64
	for angle := angle1; angle <= angle2; angle += dangle {
		x := radius * math.Sin(angle)
		y := radius * math.Cos(angle)
		point := [2]float64{x, y}
		coords = append(coords, point)
	}
	return coords
}

// Make coordinates of polygon that sits inside circle of given radius
func make_polygon_points(sides int, radius, rotate float64) [][2]float64 {
	dangle := (2.0 * math.Pi) / float64(sides)
	coords := circle_points(rotate, 2*math.Pi+rotate, dangle, radius)
	return coords
}

//func make_pointy_thing
func make_pointy_points(orbits, dangle, radius, rotate float64) [][2]float64 {
	coords := circle_points(rotate, orbits*math.Pi+rotate, dangle, radius)
	return coords
}

// Draw array of x/y coordinates starting at given x/y position to an existing pixel IMDraw
func draw_points(x, y float64, coords [][2]float64, r, g, b float64, imd *imdraw.IMDraw) {
	imd.Color = pixel.RGB(r, g, b)
	for _, point := range coords {
		imd.Push(pixel.V(x+point[0], y+point[1]))
	}
	imd.Line(1)
}

func run() {
	// Create pixelgl Window
	cfg := pixelgl.WindowConfig{
		Title:  "Join the dots",
		Bounds: pixel.R(0, 0, Width, Height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println("Hmm")
	imd := imdraw.New(nil)

	// Make some polygons
	sides := []int{3, 4, 5, 6}
	for si, s := range sides {
		polypoints := make_polygon_points(s, 80.0, 0.0)
		draw_points(Fwidth*.2, Fheight*0.25*float64(si)+80, polypoints, 1, 1, 0, imd)
	}

	// Stars
	star_shape := make_pointy_points(5.0, math.Pi*.8, 100, 0.0)
	draw_points(Fwidth*.5, Fheight*.8, star_shape, 1, 0, 1, imd)

	star_shape = make_pointy_points(5.0, math.Pi*.8, 100, 1.0)
	draw_points(Fwidth*.5, Fheight*.5, star_shape, 1, 0, 1, imd)

	star_shape = make_pointy_points(5.0, math.Pi*.8, 100, 2.0)
	draw_points(Fwidth*.5, Fheight*.2, star_shape, 1, 0, 1, imd)

	// Make some pointy things
	pointy := make_pointy_points(12.0, math.Pi/1.2, 100, 0.0)
	draw_points(Fwidth*.8, Fheight*.8, pointy, 0, 1, 1, imd)

	pointy = make_pointy_points(12.0, math.Pi/2.2, 100, 0.0)
	draw_points(Fwidth*.8, Fheight*.5, pointy, 0, 1, 1, imd)

	pointy = make_pointy_points(12.0, math.Pi/3.2, 100, 0.0)
	draw_points(Fwidth*.8, Fheight*.2, pointy, 0, 1, 1, imd)

	// Window update
	for !win.Closed() {
		// Screen clear
		win.Clear(pixel.RGB(0, 0, 0))
		//imd.Clear()
		imd.Draw(win)
		// Update the window to make changes visible
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
