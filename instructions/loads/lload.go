package loads

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

// Load long from local variable
type LLOAD struct{ base.Index8Instruction }

func (ld *LLOAD) Execute(frame *rtdz.Frame) {
	_lload(frame, uint(ld.Index))
}

type LLOAD_0 struct{ base.NoOperandsInstruction }

func (ld *LLOAD_0) Execute(frame *rtdz.Frame) {
	_lload(frame, 0)
}

type LLOAD_1 struct{ base.NoOperandsInstruction }

func (ld *LLOAD_1) Execute(frame *rtdz.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct{ base.NoOperandsInstruction }

func (ld *LLOAD_2) Execute(frame *rtdz.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct{ base.NoOperandsInstruction }

func (ld *LLOAD_3) Execute(frame *rtdz.Frame) {
	_lload(frame, 3)
}

func _lload(frame *rtdz.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
