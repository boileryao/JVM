package misc

import (
	"JVM/native"
	"JVM/rtdz"
	"JVM/rtdz/heap"
	"JVM/instructions/base"
)

// misc, abbreviation of Miscellaneous, somewhat a collection of different kind of things, "杂,繁杂,冗杂" in Chinese
// sun.misc package provides many utility classes/methods

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

func initialize(frame *rtdz.Frame) {
	vmClass := frame.Method().Class()
	savedProps := vmClass.GetRefVar("savedProps", "Ljava/util/Properties;")

	key := heap.JString(vmClass.Loader(), "foo")
	val := heap.JString(vmClass.Loader(), "bar")

	frame.OperandStack().PushRef(savedProps)
	frame.OperandStack().PushRef(key)
	frame.OperandStack().PushRef(val)

	propClass := vmClass.Loader().LoadClass("java/util/Properties")
	setPropMethod := propClass.GetInstanceMethod("setProperty",
		"(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")

	base.InvokeMethod(frame, setPropMethod)
}
