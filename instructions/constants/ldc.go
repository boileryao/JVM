package constants

import "JVM/instructions/base"
import (
	"JVM/rtdz"
	"JVM/rtdz/heap"
)

// Push item from run-time constant pool
type LDC struct{ base.Index8Instruction }

func (ldc *LDC) Execute(frame *rtdz.Frame) {
	_ldc(frame, ldc.Index)
}

// Push item from run-time constant pool (wide index)
type LDC_W struct{ base.Index16Instruction }

func (ldc *LDC_W) Execute(frame *rtdz.Frame) {
	_ldc(frame, ldc.Index)
}

func _ldc(frame *rtdz.Frame, index uint) {
	stack := frame.OperandStack()
	class := frame.Method().Class()
	c := class.ConstantPool().GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case string:
		internedStr := heap.JString(class.Loader(), c.(string))
		stack.PushRef(internedStr)
	case *heap.ClassRef: // class-object
		classRef := c.(*heap.ClassRef)
		classObj := classRef.ResolvedClass().JClass()
		stack.PushRef(classObj)
		// case MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}

// Push long or double from run-time constant pool (wide index)
type LDC2_W struct{ base.Index16Instruction }

func (ldc *LDC2_W) Execute(frame *rtdz.Frame) {
	stack := frame.OperandStack()
	pool := frame.Method().Class().ConstantPool()
	c := pool.GetConstant(ldc.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
