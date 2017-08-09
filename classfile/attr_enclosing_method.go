package classfile

/*
EnclosingMethod_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 class_index;
    u2 method_index;
}
*/
type EnclosingMethodAttribute struct {
	cp          ConstantPool
	classIndex  uint16
	methodIndex uint16
}

func (em *EnclosingMethodAttribute) readInfo(reader *ClassReader) {
	em.classIndex = reader.readUint16()
	em.methodIndex = reader.readUint16()
}

func (em *EnclosingMethodAttribute) ClassName() string {
	return em.cp.getClassName(em.classIndex)
}

func (em *EnclosingMethodAttribute) MethodNameAndDescriptor() (string, string) {
	if em.methodIndex > 0 {
		return em.cp.getNameAndType(em.methodIndex)
	} else {
		return "", ""
	}
}
