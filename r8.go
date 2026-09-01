// Package r8 emulates a simple CPU called the R8.
package r8

type CPU struct {
	PC     int
	Mem [256]int
}

func New() *CPU {
	return &CPU{}
}

func (cpu *CPU) Step() {
	// Over to you to implement `Step`!
}
