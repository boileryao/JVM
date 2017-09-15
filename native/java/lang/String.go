package lang

import "JVM/native"
import "JVM/rtdz"
import "JVM/rtdz/heap"

const jlString = "java/lang/String"

func init() {
	native.Register(jlString, "intern", "()Ljava/lang/String;", intern)
}

// public native String intern();
// ()Ljava/lang/String;
func intern(frame *rtdz.Frame) {
	this := frame.LocalVars().GetRef(0)
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
