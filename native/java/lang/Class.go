package lang

import (
	"JVM/native"
	"JVM/rtdz"
	"JVM/rtdz/heap"
)

func init() {
	native.Register("java/lang/Class", "getPrimitiveClass",
		"(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register("java/lang/Class", "getName0",
		"()Ljava/lang/String;", getName0)
	native.Register("java/lang/Class", "desiredAssertionStatus0",
		"(Ljava/lang/Class;)Z", desiredAssertionStatus0)
}

// static native Class<?> getPrimitiveClass(String name)
func getPrimitiveClass(frame *rtdz.Frame) {
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)

	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()

	frame.OperandStack().PushRef(class)
}

// private native String getName0()
func getName0(frame *rtdz.Frame) {
	this := frame.LocalVars().GetRef(0) // 'this'
	class := this.Extra().(*heap.Class)

	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)

	frame.OperandStack().PushRef(nameObj)
}

// private native boolean desiredAssertionStatus0
func desiredAssertionStatus0(frame *rtdz.Frame) {
	frame.OperandStack().PushInt(0) // push false
}
