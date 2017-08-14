package control

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

// Branch always
type GOTO struct{ base.BranchInstruction }

func (go2 *GOTO) Execute(frame *rtdz.Frame) {
	base.Branch(frame, go2.Offset)
}
