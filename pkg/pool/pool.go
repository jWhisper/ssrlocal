package pool

import (
	"math/bits"
	"sync"
)

const (
	num     = 17
	maxSize = 1 << (num - 1)
)

var (
	sizes [num]int
	pools [num]sync.Pool
)

func init() {
	for i := 0; i < num; i++ {
		size := 1 << i
		sizes[i] = size
		pools[i].New = func() interface{} {
			return make([]byte, size)
		}
	}
}

// Getbuf 1<= s <= 1<<16
func GetBuf(s int) []byte {
	if s >= 1 && s <= maxSize {
		i := bits.Len32(uint32(s)) - 1
		if sizes[i] < s {
			i++
		}
		return pools[i].Get().([]byte)[:s]
	}
	return make([]byte, s)
}

func PutBuf(b []byte) {
	if s := cap(b); s >= 1 && s <= maxSize {
		i := bits.Len32(uint32(s)) - 1
		if sizes[i] == s {
			// should do this or not:
			// b = b[:0]
			pools[i].Put(b)
		}
	}
}
