package log

// Level is logger level
type Level int8

var globalLevel = LvDebug

func init() {
	globalLevel = LvInfo
}

// logger Level
const (
	LvDebug Level = iota
	LvInfo
	LvError
)

// Active whether can print or not
func (l Level) Active() bool {
	return l >= globalLevel
}

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
