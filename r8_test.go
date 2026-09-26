package r8_test

import (
	"testing"

	"github.com/bitfield/go-r8"
)

func TestNew_InitialisesCPU(t *testing.T) {
	t.Parallel()
	cpu := r8.New()
	if cpu.PC != 0 {
		t.Errorf("after New, want pc == 0, got %d", cpu.PC)
	}
	got := cpu.Mem[0]
	if got != 0 {
		t.Errorf("after New, want Memory[0] == 0, got %d", got)
	}
}

func TestStep_IncrementsPC(t *testing.T) {
	t.Parallel()
	cpu := r8.New()
	cpu.Mem[0] = 1
	cpu.Step()
	if cpu.PC != 1 {
		t.Errorf("want pc == 1, got %d", cpu.PC)
	}
}

func TestRun_RunsUntilHalted(t *testing.T){
	t.Parallel()
	cpu := r8.New()
	cpu.RunProgram([]byte{r8.NoOp, r8.Halt})
	if cpu.PC != 2 {
		t.Errorf("want pc == 2, got %d", cpu.PC)
	}
}

func TestInc_IncrementsAccumulator(t *testing.T){
	t.Parallel()
	cpu := r8.New()
	cpu.A = 0
	cpu.RunProgram([]byte{r8.INCREMENT})
	if cpu.A != 1 {
		t.Errorf("want accumulator to be 1, got %d", cpu.A)
	}
}

func TestDec_DecrementsAccumulator(t *testing.T){
	t.Parallel()
	cpu := r8.New()
	cpu.A = 1
	cpu.RunProgram([]byte{r8.DECREMENT})
	if cpu.A != 0 {
		t.Errorf("want accumulator to be , got %d", cpu.A)
	}
}

func TestAccumulator_WrapsAroundIncrementFrom255To0(t *testing.T){
	t.Parallel()
	cpu := r8.New()
	cpu.A = 255
	cpu.RunProgram([]byte{r8.INCREMENT})
	if cpu.A != 0 {
		t.Errorf("want accumulator to be 0, got %d", cpu.A)
	}
}

func TestAccumulator_WrapsDecrementFrom0To255(t *testing.T){
	t.Parallel()
	cpu := r8.New()
	cpu.A = 0
	cpu.RunProgram([]byte{r8.DECREMENT})
	if cpu.A != 255 {
		t.Errorf("want accumulator to be 255, got %d", cpu.A)
	}
}
