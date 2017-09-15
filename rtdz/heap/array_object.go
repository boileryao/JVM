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
		panic("call len() on an non-array object!")
	}
}

func ArrayCopy(src, dst *Object, srcPos, dstPos, len int32) {
	switch src.data.(type) {
	case []int8:
		_src := src.data.([]int8)[srcPos: srcPos+len]
		_dst := dst.data.([]int8)[dstPos: dstPos+len]
		copy(_dst, _src)
	case []int16:
		_src := src.data.([]int16)[srcPos: srcPos+len]
		_dst := dst.data.([]int16)[dstPos: dstPos+len]
		copy(_dst, _src)
	case []int32:
		_src := src.data.([]int32)[srcPos: srcPos+len]
		_dst := dst.data.([]int32)[dstPos: dstPos+len]
		copy(_dst, _src)
	case []int64:
		_src := src.data.([]int64)[srcPos: srcPos+len]
		_dst := dst.data.([]int64)[dstPos: dstPos+len]
		copy(_dst, _src)
	case []uint16:
		_src := src.data.([]uint16)[srcPos: srcPos+len]
		_dst := dst.data.([]uint16)[dstPos: dstPos+len]
		copy(_dst, _src)
	case []float32:
		_src := src.data.([]float32)[srcPos: srcPos+len]
		_dst := dst.data.([]float32)[dstPos: dstPos+len]
		copy(_dst, _src)
	case []float64:
		_src := src.data.([]float64)[srcPos: srcPos+len]
		_dst := dst.data.([]float64)[dstPos: dstPos+len]
		copy(_dst, _src)
	case []*Object:
		_src := src.data.([]*Object)[srcPos: srcPos+len]
		_dst := dst.data.([]*Object)[dstPos: dstPos+len]
		copy(_dst, _src)
	default:
		panic("Not array!")
	}
}
