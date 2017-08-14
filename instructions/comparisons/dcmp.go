package comparisons

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Compare double
type DCMPG struct{ base.NoOperandsInstruction }

func (cmp *DCMPG) Execute(frame *rtdz.Frame) {
	_dcmp(frame, true)
}

type DCMPL struct{ base.NoOperandsInstruction }

func (cmp *DCMPL) Execute(frame *rtdz.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *rtdz.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
