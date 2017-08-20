package heap

// jvm spec8 6.5.instanceof
// jvm spec8 6.5.checkcast
func (kls *Class) isAssignableFrom(other *Class) bool {
	s, t := other, kls

	if s == t {
		return true
	}

	if !t.IsInterface() {
		return s.isSubClassOf(t)
	} else {
		return s.isImplements(t)
	}
}

// ref extends c
func (kls *Class) isSubClassOf(other *Class) bool {
	for c := kls.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// ref implements iface
func (kls *Class) isImplements(iface *Class) bool {
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
