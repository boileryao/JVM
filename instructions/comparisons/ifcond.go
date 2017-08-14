package comparisons

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Branch if int comparison with zero succeeds
type IFEQ struct{ base.BranchInstruction }

func (cmp *IFEQ) Execute(frame *rtdz.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, cmp.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }

func (cmp *IFNE) Execute(frame *rtdz.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, cmp.Offset)
	}
}

type IFLT struct{ base.BranchInstruction }

func (cmp *IFLT) Execute(frame *rtdz.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, cmp.Offset)
	}
}

type IFLE struct{ base.BranchInstruction }

func (cmp *IFLE) Execute(frame *rtdz.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, cmp.Offset)
	}
}

type IFGT struct{ base.BranchInstruction }

func (cmp *IFGT) Execute(frame *rtdz.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, cmp.Offset)
	}
}

type IFGE struct{ base.BranchInstruction }

func (cmp *IFGE) Execute(frame *rtdz.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, cmp.Offset)
	}
}
