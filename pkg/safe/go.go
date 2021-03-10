package safe

import (
	"fmt"
	"runtime"
)

func Go(f func()) {
	go func() {
		defer func() {
			if rerr := recover(); rerr != nil {
				fmt.Printf("Error:panic happen:%s\n%s\n", rerr, printStack())
			}
		}()
		f()
	}()
}

func printStack() string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}
