package stack

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

// Duplicate the top operand stack value
type DUP struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
             \_
               |
               V
[...][c][b][a][a]
*/
func (dup *DUP) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

// Duplicate the top operand stack value and insert two values down
type DUP_X1 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
          __/
         |
         V
[...][c][a][b][a]
*/
func (dup *DUP_X1) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Duplicate the top operand stack value and insert two or three values down
type DUP_X2 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
       _____/
      |
      V
[...][a][c][b][a]
*/
func (dup *DUP_X2) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Duplicate the top one or two operand stack values
type DUP2 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]____
          \____   |
               |  |
               V  V
[...][c][b][a][b][a]
*/
func (dup *DUP2) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Duplicate the top one or two operand stack values and insert two or three values down
type DUP2_X1 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
       _/ __/
      |  |
      V  V
[...][b][a][c][b][a]
*/
func (dup *DUP2_X1) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Duplicate the top one or two operand stack values and insert two, three, or four values down
type DUP2_X2 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][d][c][b][a]
       ____/ __/
      |   __/
      V  V
[...][b][a][d][c][b][a]
*/
func (dup *DUP2_X2) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
