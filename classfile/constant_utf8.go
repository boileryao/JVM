package classfile

import (
	"fmt"
	"unicode/utf16"
)

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
type ConstantUtf8Info struct {
	str string
}

func (string *ConstantUtf8Info) readInfo(reader *ClassReader) {

	length := uint32(reader.readUint32())
	bytes := reader.readBytes(length)
	string.str = decodeMutf8(bytes)
}

func decodeMutf8(bytes []byte) string {
	return string(bytes)
	utf_len := len(bytes)
	chars := make([]uint16, utf_len)

	var c, c2, c3 uint16
	cnt := 0
	chars_cnt := 0

	for cnt < utf_len {
		c = uint16(bytes[cnt])
		if c > 127 {
			break
		}
		cnt++
		chars[chars_cnt] = c
		chars_cnt++
	}

	for cnt < utf_len {
		c = uint16(bytes[cnt])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7: // 0xxxxxxx
			cnt++
			chars[chars_cnt] = c
			chars_cnt++
		case 12, 13: // 110xxxxx 10xxxxxx
			cnt += 2
			if cnt > utf_len {
				panic("Malformed UTF: Partial character at end")
			}
			c2 = uint16(bytes[cnt-1])
			if c2&0xC0 != 0x80 { // c2 is not 10xxxx
				panic(fmt.Errorf("Malformed UTF around byte %v", cnt))
			}
			chars[chars_cnt] = c&0x1F<<6 | c2&0x3F
			chars_cnt++
		case 14: // 1110xxxx 10xxxxxx 10xxxxxx
			cnt += 3
			if cnt > utf_len {
				panic("Malformed UTF: Partial character at end")
			}
			c2 = uint16(bytes[cnt-2])
			c3 = uint16(bytes[cnt-1])
			if c2&0xC0 != 0x80 || c3&0xC0 != 0x80 {
				panic(fmt.Errorf("Malformed UTF around byte %v", cnt))
			}
			chars[chars_cnt] = c&0x0F<<12 | c2&0x3F<<6 | c3&0x3F<<0
		}
	}
	chars = chars[0:chars_cnt]
	runes := utf16.Decode(chars)

	return string(runes) //return string(bytes) is a young,simple,naive implementation
}
