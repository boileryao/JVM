package heap

type Object struct {
	class *Class
	data  interface{}
	extra interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

func (obj *Object) Clone() *Object {
	return &Object{
		class: obj.class,
		data:  obj.cloneData(),
	}
}

// getters
func (obj *Object) Class() *Class {
	return obj.class
}
func (obj *Object) Fields() Slots {
	return obj.data.(Slots)
}
func (obj *Object) Extra() interface{} {
	return obj.extra
}

func (obj *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(obj.class)
}

func (obj *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := obj.class.getField(name, descriptor, false)
	slots := obj.data.(Slots)
	slots.SetRef(field.slotId, ref)
}
func (obj *Object) GetRefVar(name, descriptor string) *Object {
	field := obj.class.getField(name, descriptor, false)
	slots := obj.data.(Slots)
	return slots.GetRef(field.slotId)
}

func (obj *Object) cloneData() interface{} {
	switch obj.data.(type) {
	case []int8:
		elements := obj.data.([]int8)
		elements2 := make([]int8, len(elements))
		copy(elements2, elements)
		return elements2
	case []int16:
		elements := obj.data.([]int16)
		elements2 := make([]int16, len(elements))
		copy(elements2, elements)
		return elements2
	case []uint16:
		elements := obj.data.([]uint16)
		elements2 := make([]uint16, len(elements))
		copy(elements2, elements)
		return elements2
	case []int32:
		elements := obj.data.([]int32)
		elements2 := make([]int32, len(elements))
		copy(elements2, elements)
		return elements2
	case []int64:
		elements := obj.data.([]int64)
		elements2 := make([]int64, len(elements))
		copy(elements2, elements)
		return elements2
	case []float32:
		elements := obj.data.([]float32)
		elements2 := make([]float32, len(elements))
		copy(elements2, elements)
		return elements2
	case []float64:
		elements := obj.data.([]float64)
		elements2 := make([]float64, len(elements))
		copy(elements2, elements)
		return elements2
	case []*Object:
		elements := obj.data.([]*Object)
		elements2 := make([]*Object, len(elements))
		copy(elements2, elements)
		return elements2
	default: // []Slot
		slots := obj.data.(Slots)
		slots2 := newSlots(uint(len(slots)))
		copy(slots2, slots)
		return slots2
	}
}
