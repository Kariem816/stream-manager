package msgs

type MessageTag uint8

const (
	MSG_SID MessageTag = iota
	MSG_OFFER
	MSG_ANSWER
	MSG_CANDIDATE
	MSG_LEN
)

type MessageError uint8

const (
	ErrNil MessageError = iota
	ErrInvalidLen
	ErrInvalidType
)
