package extended

import (
	"JVM/instructions/base"
	"JVM/rtdz"
)

// Branch always (wide index)
type GOTO_W struct {
	offset int
}

func (ext *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	ext.offset = int(reader.ReadInt32())
}
func (ext *GOTO_W) Execute(frame *rtdz.Frame) {
	base.Branch(frame, ext.offset)
}
