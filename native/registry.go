package native

import (
	"JVM/rtdz"
	"log"
)

type NativeMathod func(frame *rtdz.Frame)

var registry = map[string]NativeMathod{}

func Register(className, methodName, descriptor string, method NativeMathod) {
	key := className + "~" + methodName + "~" + descriptor
	registry[key] = method
}

func FindNativeMethod(className, methodName, descriptor string) NativeMathod {
	key := className + "~" + methodName + "~" + descriptor
	if method, ok := registry[key]; ok{
		return method
	}

	// handle native-methods that we do not care
	if descriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}

	log.Println(key)

	return nil
}

func emptyNativeMethod(frame *rtdz.Frame) {
	// just empty
}