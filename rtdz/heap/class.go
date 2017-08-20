package heap

import "strings"
import "JVM/classfile"

// name, superClassName and interfaceNames are all binary names(jvm spec8-4.2.1)
type Class struct {
	accessFlags       uint16
	name              string // thisClassName
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (kls *Class) IsPublic() bool {
	return 0 != kls.accessFlags&ACC_PUBLIC
}
func (kls *Class) IsFinal() bool {
	return 0 != kls.accessFlags&ACC_FINAL
}
func (kls *Class) IsSuper() bool {
	return 0 != kls.accessFlags&ACC_SUPER
}
func (kls *Class) IsInterface() bool {
	return 0 != kls.accessFlags&ACC_INTERFACE
}
func (kls *Class) IsAbstract() bool {
	return 0 != kls.accessFlags&ACC_ABSTRACT
}
func (kls *Class) IsSynthetic() bool {
	return 0 != kls.accessFlags&ACC_SYNTHETIC
}
func (kls *Class) IsAnnotation() bool {
	return 0 != kls.accessFlags&ACC_ANNOTATION
}
func (kls *Class) IsEnum() bool {
	return 0 != kls.accessFlags&ACC_ENUM
}

// getters
func (kls *Class) ConstantPool() *ConstantPool {
	return kls.constantPool
}
func (kls *Class) StaticVars() Slots {
	return kls.staticVars
}

// jvm spec 5.4.4
func (kls *Class) isAccessibleTo(other *Class) bool {
	return kls.IsPublic() ||
		kls.getPackageName() == other.getPackageName()
}

func (kls *Class) getPackageName() string {
	if i := strings.LastIndex(kls.name, "/"); i >= 0 {
		return kls.name[:i]
	}
	return ""
}

func (kls *Class) GetMainMethod() *Method {
	return kls.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (kls *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range kls.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func (kls *Class) NewObject() *Object {
	return newObject(kls)
}

func (kls *Class) getMainMethod() *Method {
	return kls.getStaticMethod("main", "([Ljava/lang/String;)V")
}
