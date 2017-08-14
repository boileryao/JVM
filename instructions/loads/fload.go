package loads

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

// Load float from local variable
type FLOAD struct{ base.Index8Instruction }

func (ld *FLOAD) Execute(frame *rtdz.Frame) {
	_fload(frame, uint(ld.Index))
}

type FLOAD_0 struct{ base.NoOperandsInstruction }

func (ld *FLOAD_0) Execute(frame *rtdz.Frame) {
	_fload(frame, 0)
}

type FLOAD_1 struct{ base.NoOperandsInstruction }

func (ld *FLOAD_1) Execute(frame *rtdz.Frame) {
	_fload(frame, 1)
}

type FLOAD_2 struct{ base.NoOperandsInstruction }

func (ld *FLOAD_2) Execute(frame *rtdz.Frame) {
	_fload(frame, 2)
}

type FLOAD_3 struct{ base.NoOperandsInstruction }

func (ld *FLOAD_3) Execute(frame *rtdz.Frame) {
	_fload(frame, 3)
}

func _fload(frame *rtdz.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}
