package conversions

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Convert int to byte
type I2B struct{ base.NoOperandsInstruction }

func (cvt *I2B) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int32(int8(i))
	stack.PushInt(b)
}

// Convert int to char
type I2C struct{ base.NoOperandsInstruction }

func (cvt *I2C) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	c := int32(uint16(i))
	stack.PushInt(c)
}

// Convert int to short
type I2S struct{ base.NoOperandsInstruction }

func (cvt *I2S) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	s := int32(int16(i))
	stack.PushInt(s)
}

// Convert int to long
type I2L struct{ base.NoOperandsInstruction }

func (cvt *I2L) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	l := int64(i)
	stack.PushLong(l)
}

// Convert int to float
type I2F struct{ base.NoOperandsInstruction }

func (cvt *I2F) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	f := float32(i)
	stack.PushFloat(f)
}

// Convert int to double
type I2D struct{ base.NoOperandsInstruction }

func (cvt *I2D) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	d := float64(i)
	stack.PushDouble(d)
}
