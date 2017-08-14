package math

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Subtract double
type DSUB struct{ base.NoOperandsInstruction }

func (op *DSUB) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}

// Subtract float
type FSUB struct{ base.NoOperandsInstruction }

func (op *FSUB) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}

// Subtract int
type ISUB struct{ base.NoOperandsInstruction }

func (op *ISUB) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}

// Subtract long
type LSUB struct{ base.NoOperandsInstruction }

func (op *LSUB) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}
