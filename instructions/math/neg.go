package math

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Negate double
type DNEG struct{ base.NoOperandsInstruction }

func (op *DNEG) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

// Negate float
type FNEG struct{ base.NoOperandsInstruction }

func (op *FNEG) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

// Negate int
type INEG struct{ base.NoOperandsInstruction }

func (op *INEG) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

// Negate long
type LNEG struct{ base.NoOperandsInstruction }

func (op *LNEG) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
