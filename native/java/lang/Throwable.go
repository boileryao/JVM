package lang

import (
	"JVM/native"
	"JVM/rtdz"
	"fmt"
	"JVM/rtdz/heap"
)

type StackTraceElement struct {
	fileName   string
	className  string
	methodName string
	lineNumber int
}

func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace",
		"(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// private native Throwable fillInStackTrace(int dummy);
// (I)Ljava/lang/Throwable;
func fillInStackTrace(frame *rtdz.Frame) {
	this := frame.LocalVars().GetRef(0)
	frame.OperandStack().PushRef(this)

	traceElements := createStackTraceElements(this, frame.Thread())
	this.SetExtra(traceElements)
}

func createStackTraceElements(tObj *heap.Object, thread *rtdz.Thread) []*StackTraceElement {
	skip := distanceToObject(tObj.Class()) + 2
	frames := thread.GetFrames()[skip:]
	traceElements := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		traceElements[i] = createStackTraceElement(frame)
	}
	return traceElements
}

func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}

func createStackTraceElement(frame *rtdz.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}

func (trace *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)",
		trace.className, trace.methodName, trace.fileName, trace.lineNumber)
}
