package lang

import (
	"JVM/native"
	"JVM/rtdz"
	"unsafe"
)

func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
	native.Register("java/lang/Object", "hashCode", "()I", hashCode)
	native.Register("java/lang/Object", "clone", "()Ljava/lang/Object;", clone)
}

func getClass(frame *rtdz.Frame) {
	this := frame.LocalVars().GetRef(0) // 'this' ref
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}

func hashCode(frame *rtdz.Frame) {
	this := frame.LocalVars().GetRef(0)
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

func clone(frame *rtdz.Frame) {
	this := frame.LocalVars().GetRef(0)
	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")

	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}

	frame.OperandStack().PushRef(this.Clone())
}