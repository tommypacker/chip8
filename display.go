package chip8

func clearScreen(cpu *CPU) {
	for i := 0; i < 64; i++ {
		for j := 0; j < 32; j++ {
			cpu.screen[i][j] = 0
		}
	}
}

func drawSprite(cpu *CPU, x byte, y byte, n byte) {
	cpu.V[15] = 0
	for row := byte(0); row < n; row++ {
		rowVal := cpu.Memory[cpu.I+uint16(row)]
		for j := byte(0); j < 8; j++ {
			bitVal := (rowVal << j) & 0x80
			if bitVal > 0 {
				if cpu.screen[x+j][y+row] == 1 {
					cpu.V[15] = 1
				}
				cpu.screen[x+j][y+row] ^= 1
			}
		}
	}
}
