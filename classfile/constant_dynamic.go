package classfile

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (method *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	method.descriptorIndex = reader.readUint16()
}

type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (handle *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	handle.referenceKind = reader.readUint8()
	handle.referenceIndex = reader.readUint16()
}

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (invoke *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	invoke.bootstrapMethodAttrIndex = reader.readUint16()
	invoke.nameAndTypeIndex = reader.readUint16()

}
