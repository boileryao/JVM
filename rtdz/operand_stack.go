package rtdz

import (
	"math"
	"JVM/rtdz/heap"
)

type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	return &OperandStack{
		slots: make([]Slot, maxStack),
	}
}

func (stack *OperandStack) PushInt(val int32) {
	stack.slots[stack.size].bits = val
	stack.size++
}

func (stack *OperandStack) PopInt() int32 {
	stack.size--
	return stack.slots[stack.size].bits
}

func (stack *OperandStack) PushLong(val int64) {
	stack.slots[stack.size].bits = int32(val)
	stack.slots[stack.size+1].bits = int32(val >> 32)
	stack.size += 2
}
func (stack *OperandStack) PopLong() int64 {
	stack.size -= 2
	var val int64 = int64(stack.slots[stack.size+1].bits)
	val = val << 32
	val += int64(stack.slots[stack.size].bits)
	return val
}

func (stack *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	stack.slots[stack.size].bits = int32(bits)
	stack.size++
}
func (stack *OperandStack) PopFloat() float32 {
	stack.size--
	bits := uint32(stack.slots[stack.size].bits)
	return math.Float32frombits(bits)
}

func (stack *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	stack.PushLong(int64(bits))
}
func (stack *OperandStack) PopDouble() float64 {
	bits := stack.PopLong()
	return math.Float64frombits(uint64(bits))
}

func (stack *OperandStack) PushRef(ref *heap.Object) {
	stack.slots[stack.size].ref = ref
	stack.size++
}
func (stack *OperandStack) PopRef() *heap.Object {
	stack.size--
	return stack.slots[stack.size].ref
}

func (stack *OperandStack) PushSlot(slot Slot) {
	stack.slots[stack.size] = slot
	stack.size++
}
func (stack *OperandStack) PopSlot() Slot {
	stack.size--
	return stack.slots[stack.size]
}
