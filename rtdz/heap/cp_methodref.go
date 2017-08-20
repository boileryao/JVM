package heap

import "JVM/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodRefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (ref *MethodRef) ResolvedMethod() *Method {
	if ref.method == nil {
		ref.resolveMethodRef()
	}
	return ref.method
}

// jvm spec8 5.4.3.3
func (ref *MethodRef) resolveMethodRef() {
	//class := ref.Class()
	// todo
}
