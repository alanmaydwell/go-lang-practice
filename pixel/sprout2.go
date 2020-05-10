package main

import (
	"image"
	"os"
	"math"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

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


func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Round and round",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	pic, err := loadPicture("bigsprout_full.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	win.Clear(colornames.Firebrick)

	angle := 0.0
	d_angle := math.Pi/60

	scale := 0.1
	d_scale  := 0.01

	// This time draw sprite within the window update loop
	// also use MAT thing
	for !win.Closed() {

		win.Clear(pixel.RGB(0, 0.3, 0))

		mat := pixel.IM
		mat = mat.ScaledXY(pixel.ZV, pixel.V(scale, scale))
		mat = mat.Rotated(pixel.ZV, angle)
		mat = mat.Moved(win.Bounds().Center())
		sprite.Draw(win, mat)


		win.Update()
		angle += d_angle
		scale += d_scale
		if scale > 1.8 || scale < 0.1 {
			d_scale = -d_scale
		}

	}
}

func main() {
	pixelgl.Run(run)
}
