package stack

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Swap the top two operand stack values
type SWAP struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
          \/
          /\
         V  V
[...][c][a][b]
*/
func (swap *SWAP) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
