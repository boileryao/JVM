package math

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Multiply double
type DMUL struct{ base.NoOperandsInstruction }

func (op *DMUL) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}

// Multiply float
type FMUL struct{ base.NoOperandsInstruction }

func (op *FMUL) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

// Multiply int
type IMUL struct{ base.NoOperandsInstruction }

func (op *IMUL) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 * v2
	stack.PushInt(result)
}

// Multiply long
type LMUL struct{ base.NoOperandsInstruction }

func (op *LMUL) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}
