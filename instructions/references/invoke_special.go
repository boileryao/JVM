package references

import (
	"JVM/instructions/base"
	"JVM/rtdz"
	"JVM/rtdz/heap"
)

// instance method, especially superclass, private and instance initialization method
type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (invoke *INVOKE_SPECIAL) Execute(frame *rtdz.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(invoke.Index).(*heap.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	// check [protect] access identifier
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {

		panic("java.lang.IllegalAccessError")
	}

	// ACC_SUPER
	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>" {

		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),
			methodRef.Name(), methodRef.Descriptor())
	}

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
