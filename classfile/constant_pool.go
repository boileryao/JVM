package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {

}

func (pool ConstantPool) getConstantInfo(index uint16) ConstantInfo {

}

func (pool ConstantPool) getNameAndType(index uint16) (string, string) {

}

func (pool ConstantPool) getClassName(index uint16) string {
	return ""
}

func (pool ConstantPool) getUtf8(index uint16) string {
	return ""
}
