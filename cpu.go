package chip8

import (
	"os"
)

/*
CPU contains all components of a chip8 cpu
*/
type CPU struct {
	memory        [4096]byte
	V             [16]byte
	I             uint16
	pc            uint16
	stack         []uint16
	delayTimer    byte
	soundTimer    byte
	screen        [64][32]byte
	screenUpdated bool
}

func Reset(cpu *CPU) {
	// Reset CPU
	cpu.I = 0
	cpu.pc = 0x200
	for i := 0; i < 16; i++ {
		cpu.V[i] = 0
	}
	clearScreen(cpu)

	// Load fontset
	for i := 0; i < 80; i++ {
		cpu.memory[i] = fontset[i]
	}

	// Load program into memory
	file, err := os.Open("PONG.ch8")
	defer file.Close()
	if err != nil {
		println("ERROR: %s", err)
		return
	}

	fi, err := file.Stat()
	if err != nil {
		println("ERROR: %s", err)
		return
	}
	buffer := make([]byte, fi.Size())
	file.Read(buffer)

	for i := 0; i < len(buffer); i++ {
		cpu.memory[i+512] = buffer[i]
	}
}

func EmulateCycle(cpu *CPU) {
	opcode := uint16(cpu.memory[cpu.pc])<<8 | uint16(cpu.memory[cpu.pc+1])

	switch opcode & 0xF000 {
	case 0x0000:
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
		break
	case 0x1000:
		opcode1NNN(cpu, opcode)
		break
	case 0x2000:
		opcode2NNN(cpu, opcode)
		break
	case 0x3000:
		opcode3XNN(cpu, opcode)
		break
	case 0x4000:
		opcode4XNN(cpu, opcode)
		break
	case 0x5000:
		opcode5XY0(cpu, opcode)
		break
	case 0x6000:
		opcode6XNN(cpu, opcode)
		break
	case 0x7000:
		opcode7XNN(cpu, opcode)
		break
	case 0x8000:
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
		break
	case 0x9000:
		opcode9XY0(cpu, opcode)
		break
	case 0xA000:
		opcodeANNN(cpu, opcode)
		break
	case 0xB000:
		opcodeBNNN(cpu, opcode)
		break
	case 0xC000:
		opcodeCXNN(cpu, opcode)
		break
	case 0xD000:
		opcodeDXYN(cpu, opcode)
		break
	case 0xE000:
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
		break
	case 0xF000:
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
			break
		}
		break
	default:
		println("Unknown opcode")
	}
}
