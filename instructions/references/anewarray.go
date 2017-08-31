package references

import "JVM/instructions/base"
import "JVM/rtdz"
import "JVM/rtdz/heap"

// Create new array of reference
type ANEW_ARRAY struct{ base.Index16Instruction }

func (anew *ANEW_ARRAY) Execute(frame *rtdz.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(anew.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}
