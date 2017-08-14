package control

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

/*
lookupswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
npairs1
npairs2
npairs3
npairs4
match-offset pairs...
*/
// Access jump table by key match and jump
type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (ctrl *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	ctrl.defaultOffset = reader.ReadInt32()
	ctrl.npairs = reader.ReadInt32()
	ctrl.matchOffsets = reader.ReadInt32s(ctrl.npairs * 2)
}

func (ctrl *LOOKUP_SWITCH) Execute(frame *rtdz.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < ctrl.npairs*2; i += 2 {
		if ctrl.matchOffsets[i] == key {
			offset := ctrl.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(ctrl.defaultOffset))
}
