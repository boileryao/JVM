package classfile

import "fmt"

const (
	CONSTANT_Utf8               = 0x01
	CONSTANT_Integer            = 0x03
	CONSTANT_Float              = 0x04
	CONSTANT_Long               = 0x05
	CONSTANT_Double             = 0x06
	CONSTANT_Class              = 0x07
	CONSTANT_String             = 0x08
	CONSTANT_FieldRef           = 0x09
	CONSTANT_MethodRef          = 0x0a
	CONSTANT_InterfaceMethodRef = 0x0b
	CONSTANT_NameAndType        = 0x0c
	CONSTANT_MethodHandle       = 0x0f
	CONSTANT_MethodType         = 0x10
	CONSTANT_InvokeDynamic      = 0x12
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, pool ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, pool)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, pool ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}

	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: pool}

	case CONSTANT_Class:
		return &ConstantClassInfo{cp: pool}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_FieldRef:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{cp: pool}}
	case CONSTANT_MethodRef:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{cp: pool}}
	case CONSTANT_InterfaceMethodRef:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberRefInfo{cp: pool}}

	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		fmt.Print("Ilegal Tag: ", tag)
		panic("java.lang.ClassFormatError: Unknown Constant Pool TAG!")
	}
}
