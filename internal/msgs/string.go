package msgs

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type String string

var ErrStringTooLong = errors.New("string too long")

func (s String) Buf() ([]byte, error) {
	if len(s) > 65535 {
		return nil, ErrStringTooLong
	}
	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, uint16(len(s)))
	buf.Write([]byte(s))
	return buf.Bytes(), nil
}

func (s *String) FromBuf(b []byte) MessageError {
	if len(b) == 0 {
		return ErrInvalidLen
	}
	length := binary.BigEndian.Uint16(b)
	if len(b) < int(2+length) {
		return ErrInvalidLen
	}
	*s = String(b[2 : 2+length])
	return ErrNil
}

func (s String) String() string {
	return string(s)
}

func (s String) Len() int {
	return len(s)
}

func (s String) MarshalBinary() ([]byte, error) {
	return s.Buf()
}

func (s *String) UnmarshalBinary(data []byte) error {
	merr := s.FromBuf(data)
	if merr != ErrNil {
		return merr
	}
	return nil
}
