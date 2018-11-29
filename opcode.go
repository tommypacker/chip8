package chip8

import (
	"math/rand"
)

func randomByte() byte {
	return byte(rand.Intn(256))
}

func opcode00NN(cpu *CPU, opcode uint16) {
	switch opcode & 0x000F {
	case 0x0000:
		opcode00E0(cpu, opcode)
		break
	case 0x000E:
		opcode00EE(cpu, opcode)
		break
	default:
		println("Unknown opcode")
	}
}

func opcode8NNN(cpu *CPU, opcode uint16) {
	switch opcode & 0x000F {
	case 0x0000:
		opcode8XY0(cpu, opcode)
		break
	case 0x0001:
		opcode8XY1(cpu, opcode)
		break
	case 0x0002:
		opcode8XY2(cpu, opcode)
		break
	case 0x0003:
		opcode8XY3(cpu, opcode)
		break
	case 0x0004:
		opcode8XY4(cpu, opcode)
		break
	case 0x0005:
		opcode8XY5(cpu, opcode)
		break
	case 0x0006:
		opcode8XY6(cpu, opcode)
		break
	case 0x0007:
		opcode8XY7(cpu, opcode)
		break
	case 0x000E:
		opcode8XYE(cpu, opcode)
		break
	default:
		println("Unknown opcode")
	}
}

func opcodeENNN(cpu *CPU, opcode uint16) {
	switch opcode & 0x000F {
	case 0x000E:
		opcodeEX9E(cpu, opcode)
		break
	case 0x0001:
		opcodeEXA1(cpu, opcode)
		break
	default:
		println("Unknown opcode")
	}
}

func opcodeFNNN(cpu *CPU, opcode uint16) {
	switch opcode & 0x000F {
	case 0x0007:
		opcodeFX07(cpu, opcode)
		break
	case 0x000A:
		opcodeFX0A(cpu, opcode)
		break
	case 0x0008:
		opcodeFX18(cpu, opcode)
		break
	case 0x000E:
		opcodeFX1E(cpu, opcode)
		break
	case 0x0009:
		opcodeFX29(cpu, opcode)
		break
	case 0x0003:
		opcodeFX33(cpu, opcode)
		break
	case 0x0005:
		opcodeFNN5(cpu, opcode)
		break
	}
}

func opcodeFNN5(cpu *CPU, opcode uint16) {
	switch opcode & 0x00F0 {
	case 0x0010:
		opcodeFX15(cpu, opcode)
		break
	case 0x0050:
		opcodeFX55(cpu, opcode)
		break
	case 0x0060:
		opcodeFX65(cpu, opcode)
		break
	default:
		println("Unknown opcode")
	}
}

func opcode00E0(cpu *CPU, opcode uint16) {
	cpu.screen = [64 * 32]byte{}
	cpu.pc += 2
}

func opcode00EE(cpu *CPU, opcode uint16) {
	cpu.pc = cpu.stack[cpu.sp] + 2
	cpu.sp--
}

func opcode1NNN(cpu *CPU, opcode uint16) {
	cpu.pc = opcode & 0x0fff
}

func opcode2NNN(cpu *CPU, opcode uint16) {
	cpu.sp++
	cpu.stack[cpu.sp] = cpu.pc
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
	if cpu.V[X] > 255-cpu.V[Y] {
		cpu.V[0xF] = 1
	} else {
		cpu.V[0xF] = 0
	}
	cpu.V[X] += cpu.V[Y]
	cpu.pc += 2
}

func opcode8XY5(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	Y := (opcode & 0x00f0) >> 4
	if cpu.V[X] < cpu.V[Y] {
		cpu.V[0xF] = 0
	} else {
		cpu.V[0xF] = 1
	}
	cpu.V[X] -= cpu.V[Y]
	cpu.pc += 2
}

func opcode8XY6(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0f00) >> 8
	cpu.V[0xF] = cpu.V[x] & 0x01
	cpu.V[x] = cpu.V[x] >> 1
	cpu.pc += 2
}

func opcode8XY7(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	Y := (opcode & 0x00f0) >> 4
	if cpu.V[Y] < cpu.V[X] {
		cpu.V[0xF] = 0
	} else {
		cpu.V[0xF] = 1
	}
	cpu.V[X] = cpu.V[Y] - cpu.V[X]
	cpu.pc += 2
}

func opcode8XYE(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0f00) >> 8
	cpu.V[0xF] = (cpu.V[x] & 0x80) >> 7
	cpu.V[x] = cpu.V[x] << 1
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
	height := (opcode & 0x000f)
	x, y := cpu.V[X], cpu.V[Y]

	cpu.V[0xF] = 0
	for yline := uint16(0); yline < height; yline++ {
		pixels := cpu.memory[cpu.I+yline]
		for xline := uint16(0); xline < 8; xline++ {
			idx := (uint16(x) + xline) + (uint16(y)+yline)*64
			if idx > uint16(len(cpu.screen)) {
				continue
			}
			if (pixels & (0x80 >> xline)) != 0 {
				if cpu.screen[idx] == 1 {
					cpu.V[0xF] = 1
				}
				cpu.screen[idx] ^= 1
			}
		}
	}
	cpu.screenUpdated = true
	cpu.pc += 2
}

func opcodeEX9E(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	if cpu.key[cpu.V[X]] != 0 {
		cpu.pc += 2
		cpu.key[cpu.V[X]] = 0
	}
	cpu.pc += 2
}

func opcodeEXA1(cpu *CPU, opcode uint16) {
	X := (opcode & 0x0f00) >> 8
	if cpu.key[cpu.V[X]] == 0 {
		cpu.pc += 2
	} else {
		cpu.key[cpu.V[X]] = 0
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
	for index, k := range cpu.key {
		if k != 0 {
			cpu.V[X] = byte(index)
			cpu.pc += 2
			break
		}
	}
	cpu.key[cpu.V[X]] = 0
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
	cpu.memory[cpu.I+2] = (VX % 100) % 10
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
