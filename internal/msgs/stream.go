package msgs

import (
	"bytes"
	"encoding/binary"
)

type MsgStreamID struct {
	StreamID uint16
}

func (h MsgStreamID) Buf() ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteByte(byte(MSG_SID))
	binary.Write(&buf, binary.BigEndian, h.StreamID)
	return buf.Bytes(), nil
}

func (h *MsgStreamID) FromGeneric(g GenericMessage) MessageError {
	if g.Tag != MSG_SID {
		return ErrInvalidType
	}
	if len(g.Data) != 2 {
		return ErrInvalidLen
	}
	h.StreamID = binary.BigEndian.Uint16(g.Data)
	return ErrNil
}
