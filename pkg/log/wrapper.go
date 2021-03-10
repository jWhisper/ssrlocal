package log

import "github.com/jWhisper/ssrlocal/pkg/maths"

type Wrapper []Logger

type logOpt struct {
	md string
}

type logOptFunc func(*logOpt)

// NewWrapper logs is debug, info, err logger slice
func NewWrapper(md string, logs ...Logger) (w Wrapper) {
	//if md == "" {
	//	md = "default:"
	//}
	w = make(Wrapper, 3)
	// ls := []Level{LvDebug, LvInfo, LvError}
	dl := WithLevelAndMeta(DefaultLogger, LvDebug, md)
	il := WithLevelAndMeta(DefaultLogger, LvInfo, md)
	rl := WithLevelAndMeta(DefaultLogger, LvError, md)
	w[0] = dl
	w[1] = il
	w[2] = rl
	for i := 0; i < maths.MinInt(3, len(logs)); i++ {
		w[i] = WithLevelAndMeta(logs[i], Level(i), md)
	}
	return
}

func (w Wrapper) Debug(msg ...interface{}) {
	w[0].Print(msg...)
}
func (w Wrapper) Info(msg ...interface{}) {
	w[1].Print(msg...)
}
func (w Wrapper) Error(msg ...interface{}) {
	w[2].Print(msg...)
}
