package loads

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Load double from local variable
type DLOAD struct{ base.Index8Instruction }

func (ld *DLOAD) Execute(frame *rtdz.Frame) {
	_dload(frame, uint(ld.Index))
}

type DLOAD_0 struct{ base.NoOperandsInstruction }

func (ld *DLOAD_0) Execute(frame *rtdz.Frame) {
	_dload(frame, 0)
}

type DLOAD_1 struct{ base.NoOperandsInstruction }

func (ld *DLOAD_1) Execute(frame *rtdz.Frame) {
	_dload(frame, 1)
}

type DLOAD_2 struct{ base.NoOperandsInstruction }

func (ld *DLOAD_2) Execute(frame *rtdz.Frame) {
	_dload(frame, 2)
}

type DLOAD_3 struct{ base.NoOperandsInstruction }

func (ld *DLOAD_3) Execute(frame *rtdz.Frame) {
	_dload(frame, 3)
}

func _dload(frame *rtdz.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
