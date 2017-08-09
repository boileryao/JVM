package classfile

/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length; // must be 2
    u2 constantvalue_index;
}
*/
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (constant *ConstantValueAttribute) readInfo(reader *ClassReader) {
	constant.constantValueIndex = reader.readUint16()
}

func (constant *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return constant.constantValueIndex
}
