package comparisons

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

// Branch if reference comparison succeeds
type IF_ACMPEQ struct{ base.BranchInstruction }

func (cmp *IF_ACMPEQ) Execute(frame *rtdz.Frame) {
	if _acmp(frame) {
		base.Branch(frame, cmp.Offset)
	}
}

type IF_ACMPNE struct{ base.BranchInstruction }

func (cmp *IF_ACMPNE) Execute(frame *rtdz.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, cmp.Offset)
	}
}

func _acmp(frame *rtdz.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}
