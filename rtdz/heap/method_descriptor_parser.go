package heap

import "strings"

type MethodDescriptorParser struct {
	raw    string
	offset int
	parsed *MethodDescriptor
}

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	parser := &MethodDescriptorParser{}
	return parser.parse(descriptor)
}

func (parser *MethodDescriptorParser) parse(descriptor string) *MethodDescriptor {
	parser.raw = descriptor
	parser.parsed = &MethodDescriptor{}
	parser.startParams()
	parser.parseParamTypes()
	parser.endParams()
	parser.parseReturnType()
	parser.finish()
	return parser.parsed
}

func (parser *MethodDescriptorParser) startParams() {
	if parser.readUint8() != '(' {
		parser.causePanic()
	}
}
func (parser *MethodDescriptorParser) endParams() {
	if parser.readUint8() != ')' {
		parser.causePanic()
	}
}
func (parser *MethodDescriptorParser) finish() {
	if parser.offset != len(parser.raw) {
		parser.causePanic()
	}
}

func (parser *MethodDescriptorParser) causePanic() {
	panic("BAD descriptor: " + parser.raw)
}

func (parser *MethodDescriptorParser) readUint8() uint8 {
	b := parser.raw[parser.offset]
	parser.offset++
	return b
}
func (parser *MethodDescriptorParser) unreadUint8() {
	parser.offset--
}

func (parser *MethodDescriptorParser) parseParamTypes() {
	for {
		t := parser.parseFieldType()
		if t != "" {
			parser.parsed.addParameterType(t)
		} else {
			break
		}
	}
}

func (parser *MethodDescriptorParser) parseReturnType() {
	if parser.readUint8() == 'V' {
		parser.parsed.returnType = "V"
		return
	}

	parser.unreadUint8()
	t := parser.parseFieldType()
	if t != "" {
		parser.parsed.returnType = t
		return
	}

	parser.causePanic()
}

func (parser *MethodDescriptorParser) parseFieldType() string {
	switch parser.readUint8() {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'L':
		return parser.parseObjectType()
	case '[':
		return parser.parseArrayType()
	default:
		parser.unreadUint8()
		return ""
	}
}

func (parser *MethodDescriptorParser) parseObjectType() string {
	unread := parser.raw[parser.offset:]
	semicolonIndex := strings.IndexRune(unread, ';')
	if semicolonIndex == -1 {
		parser.causePanic()
		return ""
	} else {
		objStart := parser.offset - 1
		objEnd := parser.offset + semicolonIndex + 1
		parser.offset = objEnd
		descriptor := parser.raw[objStart:objEnd]
		return descriptor
	}
}

func (parser *MethodDescriptorParser) parseArrayType() string {
	arrStart := parser.offset - 1
	parser.parseFieldType()
	arrEnd := parser.offset
	descriptor := parser.raw[arrStart:arrEnd]
	return descriptor
}
