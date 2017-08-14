package loads

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

// Load int from local variable
type ILOAD struct{ base.Index8Instruction }

func (ld *ILOAD) Execute(frame *rtdz.Frame) {
	_iload(frame, uint(ld.Index))
}

type ILOAD_0 struct{ base.NoOperandsInstruction }

func (ld *ILOAD_0) Execute(frame *rtdz.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct{ base.NoOperandsInstruction }

func (ld *ILOAD_1) Execute(frame *rtdz.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct{ base.NoOperandsInstruction }

func (ld *ILOAD_2) Execute(frame *rtdz.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct{ base.NoOperandsInstruction }

func (ld *ILOAD_3) Execute(frame *rtdz.Frame) {
	_iload(frame, 3)
}

func _iload(frame *rtdz.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
