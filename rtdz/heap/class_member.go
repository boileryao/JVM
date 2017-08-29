package heap

import "JVM/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (member *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	member.accessFlags = memberInfo.AccessFlags()
	member.name = memberInfo.Name()
	member.descriptor = memberInfo.Descriptor()
}

func (member *ClassMember) IsPublic() bool {
	return 0 != member.accessFlags&ACC_PUBLIC
}
func (member *ClassMember) IsPrivate() bool {
	return 0 != member.accessFlags&ACC_PRIVATE
}
func (member *ClassMember) IsProtected() bool {
	return 0 != member.accessFlags&ACC_PROTECTED
}
func (member *ClassMember) IsStatic() bool {
	return 0 != member.accessFlags&ACC_STATIC
}
func (member *ClassMember) IsFinal() bool {
	return 0 != member.accessFlags&ACC_FINAL
}
func (member *ClassMember) IsSynthetic() bool {
	return 0 != member.accessFlags&ACC_SYNTHETIC
}

// getters
func (member *ClassMember) Name() string {
	return member.name
}
func (member *ClassMember) Descriptor() string {
	return member.descriptor
}
func (member *ClassMember) Class() *Class {
	return member.class
}

// jvm spec 5.4.4
func (member *ClassMember) isAccessibleTo(d *Class) bool {
	if member.IsPublic() {
		return true
	}
	c := member.class
	if member.IsProtected() {
		return d == c || d.IsSubClassOf(c) ||
			c.GetPackageName() == d.GetPackageName()
	}
	if !member.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}
	return d == c
}
