package math

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Boolean XOR int
type IXOR struct{ base.NoOperandsInstruction }

func (op *IXOR) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

// Boolean XOR long
type LXOR struct{ base.NoOperandsInstruction }

func (op *LXOR) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
