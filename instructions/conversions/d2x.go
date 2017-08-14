package conversions

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

// Convert double to float
type D2F struct{ base.NoOperandsInstruction }

func (cvt *D2F) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

// Convert double to int
type D2I struct{ base.NoOperandsInstruction }

func (cvt *D2I) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

// Convert double to long
type D2L struct{ base.NoOperandsInstruction }

func (cvt *D2L) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
