package heap

import "JVM/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldRefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (ref *FieldRef) ResolvedField() *Field {
	if ref.field == nil {
		ref.resolveFieldRef()
	}
	return ref.field
}

// jvm spec 5.4.3.2
func (ref *FieldRef) resolveFieldRef() {
	d := ref.cp.class
	c := ref.ResolvedClass()
	field := lookupField(c, ref.name, ref.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	ref.field = field
}

func lookupField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, intfs := range c.interfaces {
		if field := lookupField(intfs, name, descriptor); field != nil {
			return field
		}
	}

	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}
