package reserved

import (
	"JVM/instructions/base"
	"JVM/rtdz"
	"JVM/native"
	_ "JVM/native/java/lang"
	_ "JVM/native/sun/misc"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (invoke *INVOKE_NATIVE) Execute(frame *rtdz.Frame) {
	method := frame.Method()
	name := method.Name()
	descriptor := method.Descriptor()
	className := method.Class().Name()

	nativeMethod := native.FindNativeMethod(className, name, descriptor)
	if nativeMethod == nil {
		info := className + "." + name + descriptor
		panic("java.lang.UnsatisfiedLinkError: " + info)
	}

	nativeMethod(frame)
}