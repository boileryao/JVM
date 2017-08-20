package classfile

type MemberInfo struct {
	pool            ConstantPool
	accessFlag      uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, pool ConstantPool) []*MemberInfo {
	cnt := reader.readUint16()
	members := make([]*MemberInfo, cnt)
	for i := range members {
		members[i] = readMember(reader, pool)
	}
	return members
}

func readMember(reader *ClassReader, pool ConstantPool) *MemberInfo {
	return &MemberInfo{
		pool:            pool,
		accessFlag:      reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, pool),
	}
}

func (mi *MemberInfo) AccessFlags() uint16 {
	return mi.accessFlag
}
func (mi *MemberInfo) Name() string {
	return mi.pool.getUtf8(mi.nameIndex)
}

func (mi *MemberInfo) Descriptor() string {
	return mi.pool.getUtf8(mi.descriptorIndex)
}

func (mi *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attr := range mi.attributes {
		switch attr.(type) {
		case *CodeAttribute:
			return attr.(*CodeAttribute)
		}
	}
	return nil
}

func (mi *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range mi.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}
