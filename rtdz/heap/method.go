package heap

import "JVM/classfile"

type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}
func newMethod(class *Class, info *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(info)
	method.copyAttributes(info)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

func (m *Method) injectCodeAttribute(returnType string) {
	m.maxStack = 4
	m.maxLocals = m.argSlotCount
	switch returnType[0] {
	case 'V': // void
		m.code = []byte{0xfe, 0xb1}
	case 'D': // double
		m.code = []byte{0xfe, 0xaf}
	case 'F': // float
		m.code = []byte{0xfe, 0xae}
	case 'J': // long
		m.code = []byte{0xfe, 0xad}
	case 'L', '[': // ref, array
		m.code = []byte{0xfe, 0xb0}
	default: // integer, short, byte, char
		m.code = []byte{0xfe, 0xac}
	}
}

func (m *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		m.maxStack = codeAttr.MaxStack()
		m.maxLocals = codeAttr.MaxLocals()
		m.code = codeAttr.Code()
	}
}

func (m *Method) calcArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes {
		m.argSlotCount++
		if paramType == "J" || paramType == "D" {
			m.argSlotCount++
		}
	}
	if !m.IsStatic() {
		m.argSlotCount++ // `this` reference
	}
}

func (m *Method) IsSynchronized() bool {
	return 0 != m.accessFlags&ACC_SYNCHRONIZED
}
func (m *Method) IsBridge() bool {
	return 0 != m.accessFlags&ACC_BRIDGE
}
func (m *Method) IsVarargs() bool {
	return 0 != m.accessFlags&ACC_VARARGS
}
func (m *Method) IsNative() bool {
	return 0 != m.accessFlags&ACC_NATIVE
}
func (m *Method) IsAbstract() bool {
	return 0 != m.accessFlags&ACC_ABSTRACT
}
func (m *Method) IsStrict() bool {
	return 0 != m.accessFlags&ACC_STRICT
}

// getters
func (m *Method) MaxStack() uint {
	return m.maxStack
}
func (m *Method) MaxLocals() uint {
	return m.maxLocals
}
func (m *Method) Code() []byte {
	return m.code
}
func (m *Method) ArgSlotCount() uint {
	return m.argSlotCount
}
