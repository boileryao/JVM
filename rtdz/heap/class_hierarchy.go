package heap

// jvm spec8 6.5.instanceof
// jvm spec8 6.5.checkcast
func (kls *Class) IsAssignableFrom(other *Class) bool {
	s, t := other, kls

	if s == t {
		return true
	}

	if !t.IsInterface() {
		return s.IsSubClassOf(t)
	} else {
		return s.IsImplements(t)
	}
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
			if i == iface || i.IsSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// ref extends iface
func (kls *Class) IsSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range kls.interfaces {
		if superInterface == iface || superInterface.IsSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// c extends self
func (kls *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(kls)
}
