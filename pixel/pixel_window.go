// Example of using pixel package to create a blank window
// of fixed size. Based on example here https://github.com/faiface/pixel/wiki/Creating-a-Window
// Note I found the following two packages also needed (not mentioned in example)
// go get github.com/faiface/glhf
// go get -u -tags=gles2 github.com/go-gl/glfw/v3.3/glfw

package main


// Need both imports. Ignore the simple `import "github.com/faiface/pixel"` from start of the eg
import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)


func run() {
	// Create the window
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Window",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	// Window update loop
	for !win.Closed() {
		win.Update()
	}
}



func main() {
	pixelgl.Run(run)
}
