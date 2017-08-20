package references

import (
	"JVM/instructions/base"
	"JVM/rtdz"
	"JVM/rtdz/heap"
)

type NEW struct {
	base.Index16Instruction
}

func (new *NEW) Execute(frame *rtdz.Frame) {
	pool := frame.Method().Class().ConstantPool()
	classRef := pool.GetConstant(new.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	if class.IsInterface() || class.IsAbstract() {  //should not be initialized
		panic("java.lang.InstantiationError, Interface or Abstract Class")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}

