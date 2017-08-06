package classfile

/*
 * type definitions of numeric types
 */

// java.int, type definition, also byte and short
type ConstantIntegerInfo struct {
	val int32
}

func (integer *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	integer.val = int32(bytes)
}

// java.long, type definition
type ConstantLongInfo struct {
	val int64
}

func (long *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	long.val = int64(bytes)
}

// java.float, type definition
type ConstantFloatInfo struct {
	val float32
}

func (float *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	float.val = float32(bytes)
}

// java.double, type definition
type ConstantDoubleInfo struct {
	val float64
}

func (double *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	double.val = float64(bytes)
}
