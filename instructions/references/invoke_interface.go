package references

import (
	"JVM/instructions/base"
	"JVM/rtdz"
	"JVM/rtdz/heap"
)

// Invoke interface method
type INVOKE_INTERFACE struct {
	index uint
	// count uint8
	// zero uint8
}

func (invoke *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	invoke.index = uint(reader.ReadUint16())
	reader.ReadUint8() // count, historical problem
	reader.ReadUint8() // must be 0, historical problem
}

func (invoke *INVOKE_INTERFACE) Execute(frame *rtdz.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(invoke.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(),
		methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
