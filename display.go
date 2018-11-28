package chip8

func clearScreen(cpu *CPU) {
	for i := 0; i < 32; i++ {
		for j := 0; j < 64; j++ {
			cpu.screen[i*64+j] = 0
		}
	}
}

func drawSprite(cpu *CPU, x byte, y byte, height uint16) {
	cpu.V[0xF] = 0
	for yline := uint16(0); yline < height; yline++ {
		pixels := cpu.memory[cpu.I+yline]
		for xline := uint16(0); xline < 8; xline++ {
			idx := (uint16(x) + xline) + (uint16(y)+yline)*64
			if (pixels & (0x80 >> xline)) != 0 {
				if cpu.screen[idx] == 1 {
					cpu.V[0xF] = 1
				}
				cpu.screen[idx] ^= 1
			}
		}
	}
	cpu.ScreenUpdated = true
}
