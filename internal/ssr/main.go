package ssr

import (
	"fmt"

	"github.com/spf13/viper"
)

func Start() error {
	s := viper.Get("server")
	fmt.Println(s)
	return nil
}
