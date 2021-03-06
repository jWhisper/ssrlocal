package log

import (
	"testing"
)

func Testlogger(t *testing.T) {
	log := WithLevelAndMeta(DefaultLogger, LvInfo, "logTest")

	log.Print("test")
}
