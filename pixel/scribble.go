package main

/*
Just trying some things out here
Would be nice to be able to draw imd onto a sprite or picture
but not sure if possible.
*/

import (
	"fmt"
	"reflect"

	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	Width  = 1024
	Height = 768
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
		Title:  "Hmm?",
		Bounds: pixel.R(0, 0, Width, Height),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	pic, err := loadPicture("bigsprout_full.png")
	fmt.Println(reflect.TypeOf(pic))

	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 0, 0)
	imd.Push(pixel.V(0, 0))
	imd.Push(pixel.V(44, 100))
	imd.Line(2)

	// Making a picture directly without loading from a file
	rect := pixel.R(100, 50, 200, 300)
	piz := pixel.MakePictureData(rect)
	fmt.Println(reflect.TypeOf(piz))

	//imd.Draw(pic)
	//imd.MakePicture(pic)

	fmt.Println(reflect.TypeOf(imd))
	for !win.Closed() {
		win.Clear(pixel.RGB(1, 1, 1))
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
