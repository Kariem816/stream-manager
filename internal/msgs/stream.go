package msgs

import (
	"bytes"
	"encoding/binary"
)

type Handshake struct{}
type HandshakeAck struct {
	StreamID uint16
}

func (h Handshake) Buf() ([]byte, error) {
	return []byte{}, nil
}

func (h *Handshake) FromGeneric(g GenericMessage) MessageError {
	if g.Tag != MSG_SHK {
		return ErrInvalidType
	}
	if len(g.Data) != 0 {
		return ErrInvalidLen
	}
	return ErrNil
}

func (h HandshakeAck) Buf() ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteByte(byte(MSG_SHKACK))
	binary.Write(&buf, binary.BigEndian, h.StreamID)
	return buf.Bytes(), nil
}

func (h *HandshakeAck) FromGeneric(g GenericMessage) MessageError {
	if g.Tag != MSG_SHKACK {
		return ErrInvalidType
	}
	if len(g.Data) != 2 {
		return ErrInvalidLen
	}
	h.StreamID = binary.BigEndian.Uint16(g.Data)
	return ErrNil
}
