package heap

func (kls *Class) IsArray() bool {
	return kls.name[0] == '['
}

func (kls *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(kls.name)
	return kls.loader.LoadClass(componentClassName)
}

func (kls *Class) NewArray(count uint) *Object {
	if !kls.IsArray() {
		panic("Not array class: " + kls.name)
	}
	switch kls.Name() {
	case "[Z":
		return &Object{kls, make([]int8, count), nil}
	case "[B":
		return &Object{kls, make([]int8, count), nil}
	case "[C":
		return &Object{kls, make([]uint16, count), nil}
	case "[S":
		return &Object{kls, make([]int16, count), nil}
	case "[I":
		return &Object{kls, make([]int32, count), nil}
	case "[J":
		return &Object{kls, make([]int64, count), nil}
	case "[F":
		return &Object{kls, make([]float32, count), nil}
	case "[D":
		return &Object{kls, make([]float64, count), nil}
	default:
		return &Object{kls, make([]*Object, count), nil}
	}
}
