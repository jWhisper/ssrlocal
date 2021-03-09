package log

// Level is logger level
type Level int8

var GlobalLevel = LvInfo

// logger Level
const (
	LvDebug Level = iota
	LvInfo
	LvError
)

// Active whether can print or not
func (l Level) Active() bool {
	return l >= GlobalLevel
}

// String get LvString
func (l Level) String() string {
	switch l {
	case LvDebug:
		return "Debug"
	case LvInfo:
		return "Info"
	case LvError:
		return "Error"
	default:
		return ""
	}
}
