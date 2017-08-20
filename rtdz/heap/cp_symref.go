package heap

// symbolic reference
type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (ref *SymRef) ResolvedClass() *Class {
	if ref.class == nil {
		ref.resolveClassRef()
	}
	return ref.class
}

// jvm spec8 5.4.3.1
func (ref *SymRef) resolveClassRef() {
	d := ref.cp.class
	c := d.loader.LoadClass(ref.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	ref.class = c
}
