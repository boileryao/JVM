package loads

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Load reference from local variable
type ALOAD struct{ base.Index8Instruction }

func (ld *ALOAD) Execute(frame *rtdz.Frame) {
	_aload(frame, uint(ld.Index))
}

type ALOAD_0 struct{ base.NoOperandsInstruction }

func (ld *ALOAD_0) Execute(frame *rtdz.Frame) {
	_aload(frame, 0)
}

type ALOAD_1 struct{ base.NoOperandsInstruction }

func (ld *ALOAD_1) Execute(frame *rtdz.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct{ base.NoOperandsInstruction }

func (ld *ALOAD_2) Execute(frame *rtdz.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct{ base.NoOperandsInstruction }

func (ld *ALOAD_3) Execute(frame *rtdz.Frame) {
	_aload(frame, 3)
}

func _aload(frame *rtdz.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}
