package references

import (
	"fmt"
	"JVM/rtdz"
	"JVM/rtdz/heap"
	"JVM/instructions/base"
)

type INVOKE_VIRTUAL struct{ base.Index16Instruction }

//trick
func (invoke *INVOKE_VIRTUAL) Execute(frame *rtdz.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(invoke.Index).(*heap.MethodRef)
	if methodRef.Name() == "println" {
		stack := frame.OperandStack()
		switch methodRef.Descriptor() {
		case "(Z)V":
			fmt.Printf("%v\n", stack.PopInt() != 0)
		case "(C)V":
			fmt.Printf("%c\n", stack.PopInt())
		case "(I)V", "(B)V", "(S)V":
			fmt.Printf("%v\n", stack.PopInt())
		case "(F)V":
			fmt.Printf("%v\n", stack.PopFloat())
		case "(J)V":
			fmt.Printf("%v\n", stack.PopLong())
		case "(D)V":
			fmt.Printf("%v\n", stack.PopDouble())
		default:
			panic("println: " + methodRef.Descriptor())
		}

		stack.PopRef()
	}
}
