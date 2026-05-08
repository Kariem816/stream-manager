package msgs

import "bytes"

type Str string

func (s Str) Buf() ([]byte, error) {
	if len(s) > 255 {
		panic("string too long")
	}
	buf := bytes.Buffer{}
	buf.WriteByte(uint8(len(s)))
	buf.Write([]byte(s))
	return buf.Bytes(), nil
}

func (s *Str) FromBuf(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	length := int(b[0])
	if len(b) < 1+length {
		return nil
	}
	*s = Str(b[1 : 1+length])
	return nil
}
