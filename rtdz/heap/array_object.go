package heap

func (obj *Object) Bytes() []int8 {
	return obj.data.([]int8)
}
func (obj *Object) Shorts() []int16 {
	return obj.data.([]int16)
}
func (obj *Object) Ints() []int32 {
	return obj.data.([]int32)
}
func (obj *Object) Longs() []int64 {
	return obj.data.([]int64)
}
func (obj *Object) Chars() []uint16 {
	return obj.data.([]uint16)
}
func (obj *Object) Floats() []float32 {
	return obj.data.([]float32)
}
func (obj *Object) Doubles() []float64 {
	return obj.data.([]float64)
}
func (obj *Object) Refs() []*Object {
	return obj.data.([]*Object)
}

func (obj *Object) ArrayLength() int32 {
	switch obj.data.(type) {
	case []int8:
		return int32(len(obj.Bytes()))
	case []int16:
		return int32(len(obj.Shorts()))
	case []int32:
		return int32(len(obj.Ints()))
	case []int64:
		return int32(len(obj.Longs()))
	case []uint16:
		return int32(len(obj.Chars()))
	case []float32:
		return int32(len(obj.Floats()))
	case []float64:
		return int32(len(obj.Doubles()))
	case []*Object:
		return int32(len(obj.Refs()))
	default:
		panic("call length() on an non-array object!")
	}
}
