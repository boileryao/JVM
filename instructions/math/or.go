package math

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Boolean OR int
type IOR struct{ base.NoOperandsInstruction }

func (op *IOR) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

// Boolean OR long
type LOR struct{ base.NoOperandsInstruction }

func (op *LOR) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}
