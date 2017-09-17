package classfile

import "fmt"

type ClassFile struct {
	//magic 0xcafebabe uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool

	accessFlags uint16
	thisClass   uint16
	superClass  uint16
	interfaces  []uint16

	fields     []*MemberInfo
	methods    []*MemberInfo
	attributes []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok { // ok equals false
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	cf.constantPool = readConstantPool(reader)

	cf.accessFlags = reader.readUint16()
	cf.thisClass = reader.readUint16()
	cf.superClass = reader.readUint16()
	cf.interfaces = reader.readUint16s()
	cf.fields = readMembers(reader, cf.constantPool)
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool)
}
func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError:Bad header of class file")
	}
}

func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()
	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if cf.minorVersion == 0 {
			return
		}
	}

	msg := "java.lang.UnsupportedClassVersionError:Ver:"
	msg += string(cf.majorVersion)
	msg += "."
	msg += string(cf.minorVersion)
	panic(msg)
}

func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}
func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}
func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}
func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}
func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}
func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}
func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

func (cf *ClassFile) SuperClassName() string {
	if cf.superClass == 0 {
		return "" //java.lang.object has NO superclass
	}
	return cf.constantPool.getClassName(cf.superClass)
}
func (cf *ClassFile) InterfaceNames() []string {
	interfaces := make([]string, len(cf.interfaces))
	for i, cpIndex := range cf.interfaces {
		interfaces[i] = cf.constantPool.getClassName(cpIndex)
	}
	return interfaces
}

func (cf *ClassFile) SourceFileAttribute() *SourceFileAttribute {
	for _, attrInfo := range cf.attributes {
		switch attrInfo.(type) {
		case *SourceFileAttribute:
			return attrInfo.(*SourceFileAttribute)
		}
	}
	return nil
}
