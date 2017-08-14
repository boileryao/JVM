package stores

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Store double into local variable
type DSTORE struct{ base.Index8Instruction }

func (store *DSTORE) Execute(frame *rtdz.Frame) {
	_dstore(frame, uint(store.Index))
}

type DSTORE_0 struct{ base.NoOperandsInstruction }

func (store *DSTORE_0) Execute(frame *rtdz.Frame) {
	_dstore(frame, 0)
}

type DSTORE_1 struct{ base.NoOperandsInstruction }

func (store *DSTORE_1) Execute(frame *rtdz.Frame) {
	_dstore(frame, 1)
}

type DSTORE_2 struct{ base.NoOperandsInstruction }

func (store *DSTORE_2) Execute(frame *rtdz.Frame) {
	_dstore(frame, 2)
}

type DSTORE_3 struct{ base.NoOperandsInstruction }

func (store *DSTORE_3) Execute(frame *rtdz.Frame) {
	_dstore(frame, 3)
}

func _dstore(frame *rtdz.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}
