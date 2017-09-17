package references

import (
	"JVM/instructions/base"
	"JVM/rtdz"
	. "JVM/rtdz/heap"
	"reflect"
)

type ATHROW struct{ base.NoOperandsInstruction }

func (athrow *ATHROW) Execute(frame *rtdz.Frame) {
	exception := frame.OperandStack().PopRef()
	if exception == nil {
		panic("java.lang.NullPointerException")
	}

	thread := frame.Thread()
	if !findAndGotoExceptionHandler(thread, exception) {
		handleUncaughtException(thread, exception)
	}
}
func findAndGotoExceptionHandler(thread *rtdz.Thread, exception *Object) bool {
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC() - 1

		handlerPC := frame.Method().FindExceptionHandler(exception.Class(), pc)
		if handlerPC > 0{
			stack := frame.OperandStack()
			stack.Clear()
			stack.PushRef(exception)
			frame.SetNextPC(handlerPC)
			return true
		}

		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}

func handleUncaughtException(thread *rtdz.Thread, exception *Object) {
	thread.Clear()

	msg := GoString(exception.GetRefVar("detailMessage", "Ljava/lang/String;"))
	println(exception.Class().JavaName() + ": " + msg)

	stackInfos := reflect.ValueOf(exception.Extra())
	for i := 0; i < stackInfos.Len(); i++ {
		stackInfo := stackInfos.Index(i).Interface().(interface{
			String() string
		})
		println("\tat " + stackInfo.String())
	}

}