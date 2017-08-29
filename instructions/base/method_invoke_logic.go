package base

import (
	"JVM/rtdz"
	"JVM/rtdz/heap"
)

func InvokeMethod(invoker *rtdz.Frame, method *heap.Method) {
	thread := invoker.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invoker.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
}
