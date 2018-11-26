package chip8

import (
	"os"
)

/*
CPU contains all components of a chip8 cpu
*/
type CPU struct {
	memory        [0xFFF]byte
	V             [16]byte // Registers
	I             uint16
	pc            uint16
	stack         []uint16
	delayTimer    byte
	soundTimer    byte
	screen        [64][32]byte
	screenUpdated bool
}

func reset(cpu *CPU) {
	// Reset CPU
	cpu.I = 0
	cpu.pc = 0x200
	for i := 0; i < 16; i++ {
		cpu.V[i] = 0
	}
	clearScreen(cpu)

	// Load program into memory
	buffer := make([]byte, 0xfff)
	file, _ := os.Open("INVADERS")
	defer file.Close()
	file.Read(buffer)
	for i := 0; i < len(buffer); i++ {
		cpu.memory[i+512] = buffer[i]
	}
}

func executeLoop(cpu *CPU) {
	opcode := uint16(cpu.memory[cpu.pc])<<8 | uint16(cpu.memory[cpu.pc+1])

	switch opcode & 0xf000 {
	case 0x1000:
		opcode1NNN(cpu, opcode)
		break
	default:
		println("Unknown opcode")
	}
}
