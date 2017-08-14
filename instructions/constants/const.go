package constants

import "JVM/instructions/base"
import "JVM/rtdz"

/* A set of instructions that, containing operand in instruction */

// Push null
type ACONST_NULL struct{ base.NoOperandsInstruction }

func (null *ACONST_NULL) Execute(frame *rtdz.Frame) {
	frame.OperandStack().PushRef(nil)
}

// Push double
type DCONST_0 struct{ base.NoOperandsInstruction }

func (d0 *DCONST_0) Execute(frame *rtdz.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type DCONST_1 struct{ base.NoOperandsInstruction }

func (d1 *DCONST_1) Execute(frame *rtdz.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

// Push float
type FCONST_0 struct{ base.NoOperandsInstruction }

func (f0 *FCONST_0) Execute(frame *rtdz.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

type FCONST_1 struct{ base.NoOperandsInstruction }

func (f1 *FCONST_1) Execute(frame *rtdz.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

type FCONST_2 struct{ base.NoOperandsInstruction }

func (f2 *FCONST_2) Execute(frame *rtdz.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// Push int constant
type ICONST_M1 struct{ base.NoOperandsInstruction }

func (im1 *ICONST_M1) Execute(frame *rtdz.Frame) {
	frame.OperandStack().PushInt(-1)
}

type ICONST_0 struct{ base.NoOperandsInstruction }

func (i0 *ICONST_0) Execute(frame *rtdz.Frame) {
	frame.OperandStack().PushInt(0)
}

type ICONST_1 struct{ base.NoOperandsInstruction }

func (i1 *ICONST_1) Execute(frame *rtdz.Frame) {
	frame.OperandStack().PushInt(1)
}

type ICONST_2 struct{ base.NoOperandsInstruction }

func (i2 *ICONST_2) Execute(frame *rtdz.Frame) {
	frame.OperandStack().PushInt(2)
}

type ICONST_3 struct{ base.NoOperandsInstruction }

func (i3 *ICONST_3) Execute(frame *rtdz.Frame) {
	frame.OperandStack().PushInt(3)
}

type ICONST_4 struct{ base.NoOperandsInstruction }

func (i4 *ICONST_4) Execute(frame *rtdz.Frame) {
	frame.OperandStack().PushInt(4)
}

type ICONST_5 struct{ base.NoOperandsInstruction }

func (i5 *ICONST_5) Execute(frame *rtdz.Frame) {
	frame.OperandStack().PushInt(5)
}

// Push long constant
type LCONST_0 struct{ base.NoOperandsInstruction }

func (l0 *LCONST_0) Execute(frame *rtdz.Frame) {
	frame.OperandStack().PushLong(0)
}

type LCONST_1 struct{ base.NoOperandsInstruction }

func (l1 *LCONST_1) Execute(frame *rtdz.Frame) {
	frame.OperandStack().PushLong(1)
}
