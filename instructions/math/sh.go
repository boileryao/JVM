package math

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

// Shift left int
type ISHL struct{ base.NoOperandsInstruction }

func (op *ISHL) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

// Arithmetic shift right int
type ISHR struct{ base.NoOperandsInstruction }

func (op *ISHR) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

// Logical shift right int
type IUSHR struct{ base.NoOperandsInstruction }

func (op *IUSHR) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

// Shift left long
type LSHL struct{ base.NoOperandsInstruction }

func (op *LSHL) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

// Arithmetic shift right long
type LSHR struct{ base.NoOperandsInstruction }

func (op *LSHR) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

// Logical shift right long
type LUSHR struct{ base.NoOperandsInstruction }

func (op *LUSHR) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
