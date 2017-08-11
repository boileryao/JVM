package rtdz

import "math"

type Slot struct {
	bits int32
	ref  *Object
}

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	return make([]Slot, maxLocals)
}

func (vars LocalVars) SetInt(index uint, val int32) {
	vars[index].bits = val
}
func (vars LocalVars) GetInt(index uint) int32 {
	return vars[index].bits
}

func (vars LocalVars) SetLong(index uint, val int64) {
	vars[index].bits = int32(val)
	vars[index+1].bits = int32(val >> 32)
}
func (vars LocalVars) GetLong(index uint) int64 {
	var val int64 = int64(vars[index+1].bits)
	val = val << 32
	val += int64(vars[index].bits)
	return val
}

func (vars LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	vars[index].bits = int32(bits)
}
func (vars LocalVars) GetFloat(index uint) float32 {
	bits := uint32(vars[index].bits)
	return math.Float32frombits(bits)
}

func (vars LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	vars.SetLong(index, int64(bits))
}
func (vars LocalVars) GetDouble(index uint) float64 {
	bits := vars.GetLong(index)
	return math.Float64frombits(uint64(bits))
}

func (vars LocalVars) SetRef(index uint, ref *Object) {
	vars[index].ref = ref
}
func (vars LocalVars) GetRef(index uint) *Object {
	return vars[index].ref
}
