package classfile

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (src *SourceFileAttribute) readInfo(reader *ClassReader) {
	src.sourceFileIndex = reader.readUint16()
}

func (src *SourceFileAttribute) FileName() string {
	return src.cp.getUtf8(src.sourceFileIndex)
}
