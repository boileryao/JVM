package stores

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Store int into local variable
type ISTORE struct{ base.Index8Instruction }

func (store *ISTORE) Execute(frame *rtdz.Frame) {
	_istore(frame, uint(store.Index))
}

type ISTORE_0 struct{ base.NoOperandsInstruction }

func (store *ISTORE_0) Execute(frame *rtdz.Frame) {
	_istore(frame, 0)
}

type ISTORE_1 struct{ base.NoOperandsInstruction }

func (store *ISTORE_1) Execute(frame *rtdz.Frame) {
	_istore(frame, 1)
}

type ISTORE_2 struct{ base.NoOperandsInstruction }

func (store *ISTORE_2) Execute(frame *rtdz.Frame) {
	_istore(frame, 2)
}

type ISTORE_3 struct{ base.NoOperandsInstruction }

func (store *ISTORE_3) Execute(frame *rtdz.Frame) {
	_istore(frame, 3)
}

func _istore(frame *rtdz.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}
