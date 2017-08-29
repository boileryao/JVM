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
	d := ref.cp.class
	c := ref.ResolvedClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(c, ref.name, ref.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	ref.method = method
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
