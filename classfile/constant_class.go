package classfile

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantClassInfo struct {
	cp        *ConstantPool
	nameIndex uint16
}

func (clsInfo *ConstantClassInfo) readInfo(reader *ClassReader) {
	clsInfo.nameIndex = reader.readUint16()
}
func (clsInfo *ConstantClassInfo) Name() string {
	return clsInfo.cp.getUtf8(clsInfo.nameIndex)
}
