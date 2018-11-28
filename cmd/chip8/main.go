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
	chip8.Initialize(cpu)

	// Create window
	cfg := pixelgl.WindowConfig{
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
		chip8.EmulateCycle(cpu)

		if cpu.ScreenUpdated {
			win.Update()
		} else {
			win.UpdateInput()
		}
		handleKeyPress(cpu, win)
		print(chip8.CurKey(cpu))
	}
}

func handleKeyPress(cpu *chip8.CPU, win *pixelgl.Window) {
	if win.Pressed(pixelgl.Key1) {
		chip8.SetKey(cpu, byte(1))
	} else if win.Pressed(pixelgl.Key2) {
		chip8.SetKey(cpu, byte(2))
	} else if win.Pressed(pixelgl.Key3) {
		chip8.SetKey(cpu, byte(3))
	} else if win.Pressed(pixelgl.Key4) {
		chip8.SetKey(cpu, byte('C'))
	} else if win.Pressed(pixelgl.KeyQ) {
		chip8.SetKey(cpu, byte(4))
	} else if win.Pressed(pixelgl.KeyW) {
		chip8.SetKey(cpu, byte(5))
	} else if win.Pressed(pixelgl.KeyE) {
		chip8.SetKey(cpu, byte(6))
	} else if win.Pressed(pixelgl.KeyR) {
		chip8.SetKey(cpu, byte('D'))
	} else if win.Pressed(pixelgl.KeyA) {
		chip8.SetKey(cpu, byte(7))
	} else if win.Pressed(pixelgl.KeyS) {
		chip8.SetKey(cpu, byte(8))
	} else if win.Pressed(pixelgl.KeyD) {
		chip8.SetKey(cpu, byte(9))
	} else if win.Pressed(pixelgl.KeyF) {
		chip8.SetKey(cpu, byte('E'))
	} else if win.Pressed(pixelgl.KeyZ) {
		chip8.SetKey(cpu, byte('A'))
	} else if win.Pressed(pixelgl.KeyX) {
		chip8.SetKey(cpu, byte(0))
	} else if win.Pressed(pixelgl.KeyC) {
		chip8.SetKey(cpu, byte('B'))
	} else if win.Pressed(pixelgl.KeyV) {
		chip8.SetKey(cpu, byte('F'))
	} else {
		chip8.ResetKey(cpu)
	}
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
