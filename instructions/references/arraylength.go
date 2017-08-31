package references

import "JVM/instructions/base"
import "JVM/rtdz"

// Get length of array
type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

func (len *ARRAY_LENGTH) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
