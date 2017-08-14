package constants

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (nop *NOP) Execute(frame *rtdz.Frame) {
	//nothing
}
