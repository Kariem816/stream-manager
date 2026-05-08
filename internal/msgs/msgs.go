package msgs

type GenericMessage struct {
	Tag  MessageTag
	Data []byte
}

func Parse(buf []byte) (GenericMessage, MessageError) {
	if len(buf) < 1 {
		return GenericMessage{}, ErrInvalidLen
	}

	t := MessageTag(buf[0])
	if t >= MSG_LEN {
		return GenericMessage{}, ErrInvalidType
	}

	return GenericMessage{
		Tag:  t,
		Data: buf[1:],
	}, ErrNil
}
