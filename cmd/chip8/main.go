package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/tommypacker/chip8"
	"golang.org/x/image/colornames"
)

const (
	cyclesPerSecond   = 300
	keyRepeatDuration = time.Second / 5
)

func run() {
	ticker := time.NewTicker(time.Second / cyclesPerSecond)
	defer ticker.Stop()

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

	// Emulator loop
	drawBoard(cpu, win)
	for !win.Closed() {
		chip8.EmulateCycle(cpu)

		if chip8.DrawFlag(cpu) {
			drawBoard(cpu, win)
		} else {
			win.UpdateInput()
		}
		handleKeys(cpu, win)
		<-ticker.C
	}
}

var (
	keyMap = map[uint16]pixelgl.Button{
		0x1: pixelgl.Key1, 0x2: pixelgl.Key2, 0x3: pixelgl.Key3, 0xC: pixelgl.Key4,
		0x4: pixelgl.KeyQ, 0x5: pixelgl.KeyW, 0x6: pixelgl.KeyE, 0xD: pixelgl.KeyR,
		0x7: pixelgl.KeyA, 0x8: pixelgl.KeyS, 0x9: pixelgl.KeyD, 0xE: pixelgl.KeyF,
		0xA: pixelgl.KeyZ, 0x0: pixelgl.KeyX, 0xB: pixelgl.KeyC, 0xF: pixelgl.KeyV,
	}
	keysDown [16]*time.Ticker
)

func handleKeys(cpu *chip8.CPU, win *pixelgl.Window) {

	for index, key := range keyMap {
		if win.JustReleased(key) {
			if keysDown[index] != nil {
				keysDown[index].Stop()
				keysDown[index] = nil
			}
		} else if win.JustPressed(key) {
			if keysDown[index] == nil {
				keysDown[index] = time.NewTicker(keyRepeatDuration)
			}
			chip8.SetKey(cpu, byte(index))
		}

		if keysDown[index] == nil {
			continue
		}
		select {
		// Wait for key press to finish
		case <-keysDown[index].C:
			chip8.SetKey(cpu, byte(index))
		default:
		}

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
