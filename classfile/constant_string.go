package classfile

type ConstantStringInfo struct {
	cp          *ConstantPool
	stringIndex uint16
}

func (strInfo *ConstantStringInfo) readInfo(reader *ClassReader) {
	strInfo.stringIndex = reader.readUint16()
}
func (strInfo *ConstantStringInfo) String() string {
	return strInfo.cp.getUtf8(strInfo.stringIndex)
}
