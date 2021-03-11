package socks5

import (
	"github.com/jWhisper/ssrlocal/configs"
	"github.com/jWhisper/ssrlocal/pkg/log"
	"github.com/spf13/viper"
)

type Option func(*options)

type options struct {
	dialtimeout, readtimeout, writetimeout                                   int
	server                                                                   []string
	typeof, sp, password, method, obfs, obfs_param, protocol, protocol_param string
	logger                                                                   log.Wrapper
}

func Logger(l log.Wrapper) Option {
	return func(o *options) { o.logger = l }
}

func Server(ss []string) Option {
	return func(o *options) { o.server = ss }
}

func Sp(s string) Option {
	return func(o *options) { o.sp = s }
}

func Typeof(s string) Option {
	return func(o *options) { o.typeof = s }
}
func Pass(s string) Option {
	return func(o *options) { o.password = s }
}
func Method(s string) Option {
	return func(o *options) { o.method = s }
}
func Obfs(s string) Option {
	return func(o *options) { o.obfs = s }
}
func ObfsParam(s string) Option {
	return func(o *options) { o.obfs_param = s }
}
func Protocol(s string) Option {
	return func(o *options) { o.protocol = s }
}
func ProtocolParam(s string) Option {
	return func(o *options) { o.protocol_param = s }
}
func Dialtimeout(s int) Option {
	return func(o *options) { o.dialtimeout = s }
}
func Readtimeout(s int) Option {
	return func(o *options) { o.readtimeout = s }
}
func Writetimeout(s int) Option {
	return func(o *options) { o.writetimeout = s }
}

func getOption(c configs.Cnf) (o []Option) {
	t := Typeof(c.GetString("type"))
	p := Pass(c.GetString("password"))
	m := Method(c.GetString("method"))
	ob := Obfs(c.GetString("obfs"))
	op := ObfsParam(c.GetString("obfs_param"))
	pt := Protocol(c.GetString("protocol"))
	ptp := ProtocolParam(c.GetString("protocol_param"))
	s := Server(c.GetStringSlice("server"))
	sp := Sp(c.GetString("server_port"))
	dt := Dialtimeout(c.GetInt("dial_timeout"))
	rt := Readtimeout(c.GetInt("read_timeout"))
	wt := Writetimeout(c.GetInt("write_timeout"))
	o = []Option{t, p, m, ob, op, pt, ptp, s, sp, dt, rt, wt}
	return
}

func GetCnfOption() (o []Option) {
	return getOption(viper.GetViper())
}
