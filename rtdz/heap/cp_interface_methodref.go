package heap

import "JVM/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodRefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (ref *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if ref.method == nil {
		ref.resolveInterfaceMethodRef()
	}
	return ref.method
}

// jvm spec8 5.4.3.4
func (ref *InterfaceMethodRef) resolveInterfaceMethodRef() {
	//class := ref.ResolveClass()
	// todo
}
