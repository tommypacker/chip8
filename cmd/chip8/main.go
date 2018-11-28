package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/tommypacker/chip8"
	"golang.org/x/image/colornames"
)

func run() {
	// Load chip8
	cpu := new(chip8.CPU)
	chip8.Initialize(cpu)

	// Create window
	cfg := pixelgl.WindowConfig{
		Title:  "Chip 8",
		Bounds: pixel.R(0, 0, 512, 256),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Draw board
	drawBoard(cpu, win)
	for !win.Closed() {
		chip8.EmulateCycle(cpu)

		if cpu.ScreenUpdated {
			drawBoard(cpu, win)
			win.Update()
		} else {
			win.UpdateInput()
		}
		handleKeyPress(cpu, win)
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

func drawBoard(cpu *chip8.CPU, win *pixelgl.Window) {
	win.Clear(colornames.Black)
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 1, 1)
	screen := chip8.Screen(cpu)
	width, height := 8.0, 8.0

	for x := 0; x < 64; x++ {
		for y := 0; y < 32; y++ {
			if screen[(31-y)*64+x] == 1 {
				imd.Push(pixel.V(width*float64(x), height*float64(y)))
				imd.Push(pixel.V(width*float64(x)+width, height*float64(y)+height))
				imd.Rectangle(0)
			}
		}
	}
	imd.Draw(win)
	win.Update()
}

func main() {
	pixelgl.Run(run)
}
