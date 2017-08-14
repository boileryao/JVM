package comparisons

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Branch if int comparison succeeds
type IF_ICMPEQ struct{ base.BranchInstruction }

func (cmp *IF_ICMPEQ) Execute(frame *rtdz.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		base.Branch(frame, cmp.Offset)
	}
}

type IF_ICMPNE struct{ base.BranchInstruction }

func (cmp *IF_ICMPNE) Execute(frame *rtdz.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		base.Branch(frame, cmp.Offset)
	}
}

type IF_ICMPLT struct{ base.BranchInstruction }

func (cmp *IF_ICMPLT) Execute(frame *rtdz.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		base.Branch(frame, cmp.Offset)
	}
}

type IF_ICMPLE struct{ base.BranchInstruction }

func (cmp *IF_ICMPLE) Execute(frame *rtdz.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		base.Branch(frame, cmp.Offset)
	}
}

type IF_ICMPGT struct{ base.BranchInstruction }

func (cmp *IF_ICMPGT) Execute(frame *rtdz.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base.Branch(frame, cmp.Offset)
	}
}

type IF_ICMPGE struct{ base.BranchInstruction }

func (cmp *IF_ICMPGE) Execute(frame *rtdz.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		base.Branch(frame, cmp.Offset)
	}
}

func _icmpPop(frame *rtdz.Frame) (val1, val2 int32) {
	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}
