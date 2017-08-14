package stores

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Store long into local variable
type LSTORE struct{ base.Index8Instruction }

func (store *LSTORE) Execute(frame *rtdz.Frame) {
	_lstore(frame, uint(store.Index))
}

type LSTORE_0 struct{ base.NoOperandsInstruction }

func (store *LSTORE_0) Execute(frame *rtdz.Frame) {
	_lstore(frame, 0)
}

type LSTORE_1 struct{ base.NoOperandsInstruction }

func (store *LSTORE_1) Execute(frame *rtdz.Frame) {
	_lstore(frame, 1)
}

type LSTORE_2 struct{ base.NoOperandsInstruction }

func (store *LSTORE_2) Execute(frame *rtdz.Frame) {
	_lstore(frame, 2)
}

type LSTORE_3 struct{ base.NoOperandsInstruction }

func (store *LSTORE_3) Execute(frame *rtdz.Frame) {
	_lstore(frame, 3)
}

func _lstore(frame *rtdz.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}
