package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (pool ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if info := pool[index]; info != nil {
		return info
	}
	panic("Error: invalid index of constant pool")
}

func (pool ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := pool.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := pool.getUtf8(ntInfo.nameIndex)
	_type := pool.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (pool ConstantPool) getClassName(index uint16) string {
	info := pool.getConstantInfo(index).(*ConstantClassInfo)
	return pool.getUtf8(info.nameIndex)
}

func (pool ConstantPool) getUtf8(index uint16) string {
	return pool.getConstantInfo(index).(*ConstantUtf8Info).str
}
