package chip8

import "math/rand"

func randomByte() byte {
	return byte(rand.Intn(256))
}

func opcode00E0(cpu *CPU, opcode uint16) {
	clearScreen(cpu)
	cpu.pc += 2
}

func opcode00EE(cpu *CPU, opcode uint16) {
	l := len(cpu.stack)
	cpu.pc, cpu.stack = cpu.stack[l-1], cpu.stack[:l-1]
}

func opcode1NNN(cpu *CPU, opcode uint16) {
	cpu.pc = opcode & 0x0fff
}

func opcode2NNN(cpu *CPU, opcode uint16) {
	cpu.stack = append(cpu.stack, cpu.pc)
	cpu.pc = opcode & 0x0fff
}

func opcode3XNN(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	VX := cpu.V[X]
	if uint16(VX) == opcode&0x00ff {
		cpu.pc += 2
	}
	cpu.pc += 2
}

func opcode4XNN(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	VX := cpu.V[X]
	if uint16(VX) != opcode&0x00ff {
		cpu.pc += 2
	}
	cpu.pc += 2
}

func opcode5XY0(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	VX := cpu.V[X]
	Y := (opcode & 0x00f0) >> 4
	VY := cpu.V[Y]
	if VX == VY {
		cpu.pc += 2
	}
	cpu.pc += 2
}

func opcode6XNN(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	cpu.V[X] = byte(opcode & 0x00ff)
	cpu.pc += 2
}

func opcode7XNN(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	cpu.V[X] += byte(opcode & 0x00ff)
	cpu.pc += 2
}

func opcode8XY0(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	Y := (opcode & 0x00f0) >> 4
	cpu.V[X] = cpu.V[Y]
	cpu.pc += 2
}

func opcode8XY1(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	Y := (opcode & 0x00f0) >> 4
	cpu.V[X] = (cpu.V[X] | cpu.V[Y])
	cpu.pc += 2
}

func opcode8XY2(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	Y := (opcode & 0x00f0) >> 4
	cpu.V[X] = (cpu.V[X] & cpu.V[Y])
	cpu.pc += 2
}

func opcode8XY3(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	Y := (opcode & 0x00f0) >> 4
	cpu.V[X] = (cpu.V[X] ^ cpu.V[Y])
	cpu.pc += 2
}

func opcode8XY4(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	Y := (opcode & 0x00f0) >> 4
	total := uint16(cpu.V[X]) + uint16(cpu.V[Y])
	if total > 255 {
		cpu.V[15] = byte(1)
	} else {
		cpu.V[15] = byte(0)
	}
	cpu.V[X] = byte(total)
	cpu.pc += 2
}

func opcode8XY5(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	Y := (opcode & 0x00f0) >> 4
	if cpu.V[X] < cpu.V[Y] {
		cpu.V[15] = 0
	} else {
		cpu.V[15] = 1
	}
	cpu.V[X] -= cpu.V[Y]
	cpu.pc += 2
}

func opcode8XY6(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	cpu.V[15] = cpu.V[X] & 0x01
	cpu.V[X] = cpu.V[X] >> 1
	cpu.pc += 2
}

func opcode8XY7(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	Y := (opcode & 0x00f0) >> 4
	if cpu.V[Y] < cpu.V[X] {
		cpu.V[15] = 0
	} else {
		cpu.V[15] = 1
	}
	cpu.V[X] = cpu.V[Y] - cpu.V[X]
	cpu.pc += 2
}

func opcode8XYE(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	cpu.V[15] = cpu.V[X] & 0x80
	cpu.V[X] = cpu.V[X] << 1
	cpu.pc += 2
}

func opcode9XY0(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	Y := (opcode & 0x00f0) >> 4
	if cpu.V[X] != cpu.V[Y] {
		cpu.pc += 2
	}
	cpu.pc += 2
}

func opcodeANNN(cpu *CPU, opcode uint16) {
	cpu.I = opcode & 0x0fff
	cpu.pc += 2
}

func opcodeBNNN(cpu *CPU, opcode uint16) {
	cpu.pc = (opcode & 0x0fff) + uint16(cpu.V[0])
}

func opcodeCXNN(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	NN := byte(opcode & 0x00ff)
	cpu.V[X] = NN & randomByte()
	cpu.pc += 2
}

func opcodeDXYN(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	Y := (opcode & 0x00f0) >> 4
	N := (opcode & 0x000f)
	drawSprite(cpu, cpu.V[X], cpu.V[Y], byte(N))
	cpu.screenUpdated = true
	cpu.pc += 2
}

func opcodeEX9E(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	if keyPressed(cpu.V[X]) {
		cpu.pc += 2
	}
	cpu.pc += 2
}

func opcodeEXA1(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	if !keyPressed(cpu.V[X]) {
		cpu.pc += 2
	}
	cpu.pc += 2
}

func opcodeFX07(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	cpu.V[X] = cpu.delayTimer
	cpu.pc += 2
}

func opcodeFX0A(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	cpu.V[X] = getKeyPressed()
	cpu.pc += 2
}

func opcodeFX15(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	cpu.delayTimer = cpu.V[X]
	cpu.pc += 2
}

func opcodeFX18(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	cpu.soundTimer = cpu.V[X]
	cpu.pc += 2
}

func opcodeFX1E(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	cpu.I += uint16(cpu.V[X])
	cpu.pc += 2
}

func opcodeFX29(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	cpu.I = uint16(cpu.V[X]) * uint16(5)
	cpu.pc += 2
}

func opcodeFX33(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	VX := cpu.V[X]
	cpu.memory[cpu.I] = VX / 100
	cpu.memory[cpu.I+1] = (VX / 10) % 10
	cpu.memory[cpu.I+2] = VX % 10
	cpu.pc += 2
}

func opcodeFX55(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	pointer := cpu.I
	for i := uint16(0); i <= X; i++ {
		cpu.memory[pointer+i] = cpu.V[i]
	}
	cpu.pc += 2
}

func opcodeFX65(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	pointer := cpu.I
	for i := uint16(0); i <= X; i++ {
		cpu.V[i] = cpu.memory[pointer+i]
	}
	cpu.pc += 2
}
