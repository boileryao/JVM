package stack

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

// Pop the top operand stack value
type POP struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
            |
            V
[...][c][b]
*/
func (pop *POP) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// Pop the top one or two operand stack values
type POP2 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
         |  |
         V  V
[...][c]
*/
func (pop *POP2) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
