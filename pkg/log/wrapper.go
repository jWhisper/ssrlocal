package log

import "github.com/jWhisper/ssrlocal/pkg/maths"

type wrapper []Logger

type logOpt struct {
	md string
}

type logOptFunc func(*logOpt)

// NewWrapper logs is debug, info, err logger slice
func NewWrapper(md string, logs ...Logger) (w wrapper) {
	if md == "" {
		md = "default:"
	}
	w = make(wrapper, 3)
	// ls := []Level{LvDebug, LvInfo, LvError}
	dl := WithLevelAndMeta(DefaultLogger, LvDebug, md)
	il := WithLevelAndMeta(DefaultLogger, LvInfo, md)
	rl := WithLevelAndMeta(DefaultLogger, LvError, md)
	w[0] = dl
	w[1] = il
	w[2] = rl
	for i := 0; i < maths.MinInt(3, len(logs)); i++ {
		w[i] = WithLevelAndMeta(logs[i], 0, md)
	}
	return
}

func (w wrapper) Debug(msg ...interface{}) {
	w[0].Print(msg...)
}
func (w wrapper) Info(msg ...interface{}) {
	w[1].Print(msg...)
}
func (w wrapper) Error(msg ...interface{}) {
	w[2].Print(msg...)
}
