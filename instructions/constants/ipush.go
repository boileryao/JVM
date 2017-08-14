package constants

import (
	"JVM/rtdz"
	"JVM/instructions/base"
)

// Push byte
type BIPUSH struct {
	val int8
}

func (bi *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	bi.val = reader.ReadInt8()
}
func (bi *BIPUSH) Execute(frame *rtdz.Frame) {
	i := int32(bi.val)
	frame.OperandStack().PushInt(i)
}

// Push short
type SIPUSH struct {
	val int16
}

func (si *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	si.val = reader.ReadInt16()
}
func (si *SIPUSH) Execute(frame *rtdz.Frame) {
	i := int32(si.val)
	frame.OperandStack().PushInt(i)
}
