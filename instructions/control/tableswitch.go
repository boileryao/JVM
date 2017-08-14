package control

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

/*
tableswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
lowbyte1
lowbyte2
lowbyte3
lowbyte4
highbyte1
highbyte2
highbyte3
highbyte4
jump offsets...
*/
// Access jump table by index and jump
type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (ctrl *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	ctrl.defaultOffset = reader.ReadInt32()
	ctrl.low = reader.ReadInt32()
	ctrl.high = reader.ReadInt32()
	jumpOffsetsCount := ctrl.high - ctrl.low + 1
	ctrl.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (ctrl *TABLE_SWITCH) Execute(frame *rtdz.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= ctrl.low && index <= ctrl.high {
		offset = int(ctrl.jumpOffsets[index-ctrl.low])
	} else {
		offset = int(ctrl.defaultOffset)
	}

	base.Branch(frame, offset)
}
