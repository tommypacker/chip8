package main

import (
	"image"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/tommypacker/chip8"
)

func run() {
	// Load chip8
	cpu := new(chip8.CPU)
	chip8.Reset(cpu)
	chip8.EmulateCycle(cpu)

	// Create window
	/*cfg := pixelgl.WindowConfig{
		Title:  "Chip 8",
		Bounds: pixel.R(0, 0, 256, 128),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Draw board
	img := image.NewGray(image.Rect(0, 0, 256, 128))
	p, _ := drawBoard(img)
	s := pixel.NewSprite(p, p.Bounds())
	s.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	for !win.Closed() {
		if win.Pressed(pixelgl.Key1) {

		} else if win.Pressed(pixelgl.Key2) {

		} else if win.Pressed(pixelgl.Key3) {

		} else if win.Pressed(pixelgl.Key4) {

		}
		win.Update()
	}*/
}

func drawBoard(img *image.Gray) (pixel.Picture, error) {
	pixels := make([]byte, 256*128)
	for i := 0; i < 256; i++ {
		for j := 0; j < 128; j++ {
			pixels[i*128+j] = 128
		}
	}
	img.Pix = pixels
	return pixel.PictureDataFromImage(img), nil
}

func main() {
	pixelgl.Run(run)
}
