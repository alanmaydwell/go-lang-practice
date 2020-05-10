package main

import (
	"image"
	"os"
	"math"
    "math/rand"
    "fmt"
	_ "image/png"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)


// Copied from https://github.com/faiface/pixel/wiki/Drawing-a-Sprite
func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

// Using floats for x/y coords because need to make pixel.Vec (pixel.V?) which requires floats
type Movable struct {
    x float64
    y float64
    dx float64
    dy float64
    a float64
    da float64
    scale  float64
    dscale float64
}


func run() {
	fmt.Println("Starting")	
	// Window configuration and creation - VSync looks handy for updates/animation
	cfg := pixelgl.WindowConfig{
		Title:  "Sproutastic",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Load picture from file
	pic, err := loadPicture("bigsprout_full.png")
	if err != nil {
		panic(err)
	}
    // Create a pixel sprite from the picture- something that we can draw in window
	sprite := pixel.NewSprite(pic, pic.Bounds())


	// Define some moving things. Create instances of Movable struct and store them in a "slice"
    movables := []Movable {} 
    for i:=0; i<12; i++ {
        // To append need myslice = append(myslice, newthing)
        newthing := Movable{
	        x: rand.Float64()*1024,
	        y: rand.Float64()*768,
	        dx: 8*(rand.Float64()*2 -1),
	        dy: 0.0,
	        a: rand.Float64()*math.Pi*2,
	        da: math.Pi/60,
	        scale: 0.2 + (0.4*rand.Float64()),
	        dscale: 0.01,
	    	}
        movables = append(movables, newthing)
    }


    // Updating the window
	for !win.Closed() {
        // Clear window contents with specified background colour
		win.Clear(pixel.RGB(0, 0.3, 0))

        // plot the multiple sprites
        // Can iterate over slice using range function - returns index no, element
        // Note element is local copy only - ok for reading but updating will not affect parent)
        for mindex, movable := range movables {
        	// Image drawing mat thing
		    mat := pixel.IM
		    mat = mat.ScaledXY(pixel.ZV, pixel.V(movable.scale, movable.scale))
		    mat = mat.Rotated(pixel.ZV, movable.a)
            // need pixel Vector pixel.V here for coordinates
		    mat = mat.Moved(pixel.V(movable.x, movable.y))
		    sprite.Draw(win, mat)

            // Update sprite values. Beware trap - to make updates need to access field by index  
            // `movable.a += movable.da` doesn't update the value in the slice - it's only local

            // Movement
            movables[mindex].x += movable.dx
            movables[mindex].y += movable.dy
            // Simple wrapping
            if movables[mindex].x < 0 {movables[mindex].x = 1024}
            if movables[mindex].x > 1024 {movables[mindex].x = 0}
            // Rotation
            movables[mindex].a += movable.da
            // Update scale
            movables[mindex].scale += movable.dscale
            // Enforce max/min limits on scale
            if movable.scale > 0.6 {
            	movables[mindex].scale = 0.6
            	movables[mindex].dscale = -0.01
            	}
            if movable.scale < 0.1{
            	movables[mindex].scale = 0.1
            	movables[mindex].dscale = 0.01
            	}
            }

        // Update the displayed details (presumably some sort of buffer flip)
		win.Update()

	}
}

func main() {
	pixelgl.Run(run)
}
