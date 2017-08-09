package classfile

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	AttributeTable
}

func (code *CodeAttribute) readInfo(reader *ClassReader) {
	code.maxStack = reader.readUint16()
	code.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	code.code = reader.readBytes(codeLength)
	code.exceptionTable = readExceptionTable(reader)
	code.attributes = readAttributes(reader, code.cp)
}

func (code *CodeAttribute) MaxStack() uint {
	return uint(code.maxStack)
}
func (code *CodeAttribute) MaxLocals() uint {
	return uint(code.maxLocals)
}
func (code *CodeAttribute) Code() []byte {
	return code.code
}
func (code *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return code.exceptionTable
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

func (self *ExceptionTableEntry) StartPc() uint16 {
	return self.startPc
}
func (self *ExceptionTableEntry) EndPc() uint16 {
	return self.endPc
}
func (self *ExceptionTableEntry) HandlerPc() uint16 {
	return self.handlerPc
}
func (self *ExceptionTableEntry) CatchType() uint16 {
	return self.catchType
}
