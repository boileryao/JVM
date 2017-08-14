package conversions

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

// Convert float to double
type F2D struct{ base.NoOperandsInstruction }

func (cvt *F2D) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

// Convert float to int
type F2I struct{ base.NoOperandsInstruction }

func (cvt *F2I) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

// Convert float to long
type F2L struct{ base.NoOperandsInstruction }

func (cvt *F2L) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	l := int64(f)
	stack.PushLong(l)
}
