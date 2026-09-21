// Package r8 emulates a simple CPU called the R8.
package r8

type CPU struct {
	A   byte
	PC  uint16
	Mem [65536]byte
}

const NoOp = 1
const Halt = 0
const INCREMENT = 48
const DECREMENT = 64

func New() *CPU {
	return &CPU{}
}

func (cpu *CPU) RunProgram(prog []byte){
	copy(cpu.Mem[:], prog)
	cpu.Run()
}

func (cpu *CPU) Step() bool {
	// Over to you to implement `Step`!
	switch cpu.Mem[cpu.PC] {
	case Halt:
		cpu.PC++
		return false
	case INCREMENT:
		cpu.A++
	case DECREMENT:
		cpu.A--
	}
	cpu.PC++
	return true
}

func (cpu *CPU) Run() {
	for cpu.Step() {
	}
}
