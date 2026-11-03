// Package r8 emulates a simple CPU called the R8.
package r8

import "fmt"

type CPU struct {
	A   byte
	PC  uint16
	Mem [65536]byte
}

const OpNOP = 1
const OpHALT = 0
const OpINCREMENT = 48
const OpDECREMENT = 64

func New() *CPU {
	return &CPU{}
}

func (cpu *CPU) RunProgram(prog []byte) {
	copy(cpu.Mem[:], prog)
	cpu.Run()
}

func (cpu *CPU) Step() bool {
	switch cpu.Mem[cpu.PC] {
	case OpHALT:
		cpu.PC++
		return false
	case OpINCREMENT:
		cpu.A++
	case OpDECREMENT:
		cpu.A--
	}
	cpu.PC++
	return true
}

func (cpu *CPU) Run() {
	for cpu.Step() {
	}
}

func (cpu *CPU) State() string {
	return fmt.Sprintf("PC=%v A=%v OP=%v", cpu.PC, cpu.A, cpu.Mem[cpu.PC])
}
