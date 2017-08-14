package stores

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Store reference into local variable
type ASTORE struct{ base.Index8Instruction }

func (store *ASTORE) Execute(frame *rtdz.Frame) {
	_astore(frame, uint(store.Index))
}

type ASTORE_0 struct{ base.NoOperandsInstruction }

func (store *ASTORE_0) Execute(frame *rtdz.Frame) {
	_astore(frame, 0)
}

type ASTORE_1 struct{ base.NoOperandsInstruction }

func (store *ASTORE_1) Execute(frame *rtdz.Frame) {
	_astore(frame, 1)
}

type ASTORE_2 struct{ base.NoOperandsInstruction }

func (store *ASTORE_2) Execute(frame *rtdz.Frame) {
	_astore(frame, 2)
}

type ASTORE_3 struct{ base.NoOperandsInstruction }

func (store *ASTORE_3) Execute(frame *rtdz.Frame) {
	_astore(frame, 3)
}

func _astore(frame *rtdz.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}
