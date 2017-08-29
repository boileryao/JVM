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
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()
	}
	return methods
}

func (m *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		m.maxStack = codeAttr.MaxStack()
		m.maxLocals = codeAttr.MaxLocals()
		m.code = codeAttr.Code()
	}
}

func (m *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(m.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
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
