package extended

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

// Branch if reference is null
type IFNULL struct{ base.BranchInstruction }

func (ext *IFNULL) Execute(frame *rtdz.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, ext.Offset)
	}
}

// Branch if reference not null
type IFNONNULL struct{ base.BranchInstruction }

func (ext *IFNONNULL) Execute(frame *rtdz.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, ext.Offset)
	}
}
