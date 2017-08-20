package references

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

//trick
func (invoke *INVOKE_SPECIAL) Execute(frame *rtdz.Frame) {
	frame.OperandStack().PopRef()
}
