package chip8

import (
	"os"
)

// NullKey is a byte that signals when no key is pressed
var NullKey = byte(0)

/*
CPU contains all components of a chip8 cpu
*/
type CPU struct {
	memory        [4096]byte
	V             [16]byte
	I             uint16
	pc            uint16
	stack         [16]uint16
	sp            byte
	delayTimer    byte
	soundTimer    byte
	screen        [64 * 32]byte
	screenUpdated bool
	key           [16]byte
}

func Initialize(cpu *CPU) {
	cpu.I = 0
	cpu.pc = 0x200
	cpu.sp = 0
	cpu.stack = [16]uint16{}
	cpu.V = [16]byte{}
	cpu.screen = [64 * 32]byte{}
	for i := 0; i < 16; i++ {
		cpu.V[i] = 0
	}

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

func SetKey(cpu *CPU, index byte) {
	cpu.key[index] = 1
}

func Screen(cpu *CPU) [64 * 32]byte {
	return cpu.screen
}

func DrawFlag(cpu *CPU) bool {
	flag := cpu.screenUpdated
	cpu.screenUpdated = false
	return flag
}

func EmulateCycle(cpu *CPU) {
	opcode := uint16(cpu.memory[cpu.pc])<<8 | uint16(cpu.memory[cpu.pc+1])

	switch opcode & 0xF000 {
	case 0x0000:
		opcode00NN(cpu, opcode)
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
		opcode8NNN(cpu, opcode)
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
		opcodeENNN(cpu, opcode)
		break
	case 0xF000:
		opcodeFNNN(cpu, opcode)
		break
	default:
		println("Unknown opcode")
	}

	if cpu.delayTimer > 0 {
		cpu.delayTimer--
	}
}
