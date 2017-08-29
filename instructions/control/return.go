package control

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

// Return void from method
type RETURN struct{ base.NoOperandsInstruction }

func (self *RETURN) Execute(frame *rtdz.Frame) {
	frame.Thread().PopFrame()
}

// Return reference
type ARETURN struct{ base.NoOperandsInstruction }

func (self *ARETURN) Execute(frame *rtdz.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(ref)
}

// Return double
type DRETURN struct{ base.NoOperandsInstruction }

func (self *DRETURN) Execute(frame *rtdz.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}

// Return float
type FRETURN struct{ base.NoOperandsInstruction }

func (self *FRETURN) Execute(frame *rtdz.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}

// Return integer
type IRETURN struct{ base.NoOperandsInstruction }

func (self *IRETURN) Execute(frame *rtdz.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	val := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(val)
}

// Return double
type LRETURN struct{ base.NoOperandsInstruction }

func (self *LRETURN) Execute(frame *rtdz.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}
