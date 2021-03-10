package proxy

import (
	"github.com/jWhisper/ssrlocal/configs"
	"github.com/spf13/viper"
)

type Option func(*options)

type options struct {
	timeout int
	lp      string
}

func Lp(s string) Option {
	return func(o *options) { o.lp = s }
}

func Timeout(s int) Option {
	return func(o *options) { o.timeout = s }
}

func getOption(c configs.Cnf) (o []Option) {
	lp := Lp(c.GetString("local_port"))
	ti := Timeout(c.GetInt("timeout"))
	o = []Option{lp, ti}
	return
}

func GetCnfOption() (o []Option) {
	return getOption(viper.GetViper())
}
