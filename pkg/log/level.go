package log

// Level is logger level
type Level int8

const (
	LvDebug Level = iota
	LvInfo
	LvError
)

// String get LvString
func (l Level) String() string {
	switch l {
	case LvDebug:
		return "DeBug"
	case LvInfo:
		return "Info"
	case LvError:
		return "Error"
	default:
		return ""
	}
}
