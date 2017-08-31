package references

import "JVM/instructions/base"
import "JVM/rtdz"
import "JVM/rtdz/heap"

// Create new multidimensional array
type MULTI_ANEW_ARRAY struct {
	index      uint16
	dimensions uint8
}

func (new *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	new.index = reader.ReadUint16()
	new.dimensions = reader.ReadUint8()
}
func (new *MULTI_ANEW_ARRAY) Execute(frame *rtdz.Frame) {
	pool := frame.Method().Class().ConstantPool()
	classRef := pool.GetConstant(uint(new.index)).(*heap.ClassRef)
	arrClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	counts := popAndCheckCounts(stack, int(new.dimensions))
	arr := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(arr)
}

func popAndCheckCounts(stack *rtdz.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}

	return counts
}

func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0])
	arr := arrClass.NewArray(count)

	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass())
		}
	}

	return arr
}
