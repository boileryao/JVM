package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (unparsed *UnparsedAttribute) readInfo(reader *ClassReader) {
	unparsed.info = reader.readBytes(unparsed.length)
}

func (unparsed *UnparsedAttribute) Info() []byte {
	return unparsed.info
}
