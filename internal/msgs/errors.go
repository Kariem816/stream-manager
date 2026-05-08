package msgs

type MessageError uint8

const (
	ErrNil MessageError = iota
	ErrInvalidLen
	ErrInvalidType
)

func (e MessageError) Error() string {
	switch e {
	case ErrNil:
		return "nil"
	case ErrInvalidLen:
		return "invalid length"
	case ErrInvalidType:
		return "invalid type"
	default:
		return "unknown error"
	}
}
