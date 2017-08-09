package classfile

/*
Signature_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 signature_index;
}
*/
type SignatureAttribute struct {
	cp             ConstantPool
	signatureIndex uint16
}

func (sign *SignatureAttribute) readInfo(reader *ClassReader) {
	sign.signatureIndex = reader.readUint16()
}

func (sign *SignatureAttribute) Signature() string {
	return sign.cp.getUtf8(sign.signatureIndex)
}
