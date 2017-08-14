package base

import "JVM/rtdz"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtdz.Frame)
}

//NoOperandsInstruction, has no operand
type NoOperandsInstruction struct {
	// empty
}

func (nop *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}


//BranchInstruction, jump the offset
type BranchInstruction struct {
	Offset int // todo target
}

func (branch *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	branch.Offset = int(reader.ReadInt16())
}


//Index8Instruction, handle single byte operand
type Index8Instruction struct {
	Index uint
}

func (index8 *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	index8.Index = uint(reader.ReadUint8())
}

//Index16Instruction, handle two bytes operand
type Index16Instruction struct {
	Index uint
}

func (index16 *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	index16.Index = uint(reader.ReadUint16())
}
