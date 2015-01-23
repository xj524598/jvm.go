package classfile

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantClassInfo struct {
    nameIndex uint16
}
func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
    self.nameIndex = reader.readUint16()
}
func (self *ConstantClassInfo) GetName(cp *ConstantPool) (string) {
    return cp.getUtf8(self.nameIndex)
}

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantMemberrefInfo struct {
    classIndex       uint16
    nameAndTypeIndex uint16
}
func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
    self.classIndex = reader.readUint16()
    self.nameAndTypeIndex = reader.readUint16()
}

type ConstantFieldrefInfo struct {
    ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
    ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
    ConstantMemberrefInfo
}
