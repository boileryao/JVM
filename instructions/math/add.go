package math

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Add double
type DADD struct{ base.NoOperandsInstruction }

func (op *DADD) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}

// Add float
type FADD struct{ base.NoOperandsInstruction }

func (op *FADD) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

// Add int
type IADD struct{ base.NoOperandsInstruction }

func (op *IADD) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

// Add long
type LADD struct{ base.NoOperandsInstruction }

func (op *LADD) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}
