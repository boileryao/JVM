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
	inited            bool
	jClass            *Object // class object
	sourceFile        string
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
	class.sourceFile = getSourceFile(cf)
	return class
}

func getSourceFile(cf *classfile.ClassFile) string {
	if attr := cf.SourceFileAttribute(); attr != nil {
		return attr.FileName()
	}
	return "UnknownSourceFileName"
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
func (kls *Class) SuperClass() *Class {
	return kls.superClass
}
func (kls *Class) Name() string {
	return kls.name
}
func (kls *Class) JClass() *Object {
	return kls.jClass
}

func (kls *Class) Inited() bool {
	return kls.inited
}
func (kls *Class) Loader() *ClassLoader {
	return kls.loader
}
func (kls *Class) SourceFile() string {
	return kls.sourceFile
}

func (kls *Class) SetInited() {
	kls.inited = true
}

func (kls *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[kls.name]
	return ok
}

// array related
func (kls *Class) ArrayClass() *Class {
	arrKlsName := getArrayClassName(kls.name)
	return kls.loader.LoadClass(arrKlsName)
}

func (kls *Class) JavaName() string {
	return strings.Replace(kls.name, "/", ".", -1)
}

// jvm spec 5.4.4
func (kls *Class) isAccessibleTo(other *Class) bool {
	return kls.IsPublic() ||
		kls.GetPackageName() == other.GetPackageName()
}

func (kls *Class) GetPackageName() string {
	if i := strings.LastIndex(kls.name, "/"); i >= 0 {
		return kls.name[:i]
	}
	return ""
}

func (kls *Class) GetMainMethod() *Method {
	return kls.getMethod("main", "([Ljava/lang/String;)V", true)
}

func (kls *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for _, method := range kls.methods {
		if method.IsStatic() == isStatic &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func (kls *Class) GetClinitMethod() *Method {
	return kls.getMethod("<clinit>", "()V", true)
}

func (kls *Class) NewObject() *Object {
	return newObject(kls)
}

func (kls *Class) GetInstanceMethod(name, descriptor string) *Method {
	return kls.getMethod(name, descriptor, false)
}

func (kls *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := kls.getField(fieldName, fieldDescriptor, true)
	return kls.staticVars.GetRef(field.slotId)
}
func (kls *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := kls.getField(fieldName, fieldDescriptor, true)
	kls.staticVars.SetRef(field.slotId, ref)
}

func (kls *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := kls; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic &&
				field.name == name &&
				field.descriptor == descriptor {

				return field
			}
		}
	}
	return nil
}
