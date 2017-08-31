package heap

// jvm spec8 6.5.instanceof
// jvm spec8 6.5.checkcast
func (kls *Class) isAssignableFrom(other *Class) bool {
	s, t := other, kls

	if s == t {
		return true
	}

	if !s.IsArray() {
		if !s.IsInterface() {
			// s is class
			if !t.IsInterface() {
				// t is not interface
				return s.IsSubClassOf(t)
			} else {
				// t is interface
				return s.IsImplements(t)
			}
		} else {
			// s is interface
			if !t.IsInterface() {
				// t is not interface
				return t.isJlObject()
			} else {
				// t is interface
				return t.isSuperInterfaceOf(s)
			}
		}
	} else {
		// s is array
		if !t.IsArray() {
			if !t.IsInterface() {
				// t is class
				return t.isJlObject()
			} else {
				// t is interface
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			// t is array
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.isAssignableFrom(sc)
		}
	}

	return false
}

// ref extends c
func (kls *Class) IsSubClassOf(other *Class) bool {
	for c := kls.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// ref implements iface
func (kls *Class) IsImplements(iface *Class) bool {
	for c := kls; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// ref extends iface
func (kls *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range kls.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// c extends kls
func (kls *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(kls)
}

// iface extends kls
func (self *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(self)
}

func (kls *Class) isJlObject() bool {
	return kls.name == "java/lang/Object"
}
func (kls *Class) isJlCloneable() bool {
	return kls.name == "java/lang/Cloneable"
}
func (kls *Class) isJioSerializable() bool {
	return kls.name == "java/io/Serializable"
}
