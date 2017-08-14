package stores

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Store float into local variable
type FSTORE struct{ base.Index8Instruction }

func (store *FSTORE) Execute(frame *rtdz.Frame) {
	_fstore(frame, uint(store.Index))
}

type FSTORE_0 struct{ base.NoOperandsInstruction }

func (store *FSTORE_0) Execute(frame *rtdz.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct{ base.NoOperandsInstruction }

func (store *FSTORE_1) Execute(frame *rtdz.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct{ base.NoOperandsInstruction }

func (store *FSTORE_2) Execute(frame *rtdz.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct{ base.NoOperandsInstruction }

func (store *FSTORE_3) Execute(frame *rtdz.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *rtdz.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
