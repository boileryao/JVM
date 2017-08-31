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
		return &Object{kls, make([]int8, count)}
	case "[B":
		return &Object{kls, make([]int8, count)}
	case "[C":
		return &Object{kls, make([]uint16, count)}
	case "[S":
		return &Object{kls, make([]int16, count)}
	case "[I":
		return &Object{kls, make([]int32, count)}
	case "[J":
		return &Object{kls, make([]int64, count)}
	case "[F":
		return &Object{kls, make([]float32, count)}
	case "[D":
		return &Object{kls, make([]float64, count)}
	default:
		return &Object{kls, make([]*Object, count)}
	}
}
