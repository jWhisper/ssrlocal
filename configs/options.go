package configs

type Option func(*Options)

type Options struct {
	Timeout                                                            int
	Type, Password, Method, Obfs, Obfs_param, Protocol, Protocol_param string
	Da                                                                 []byte
}

func Type(s string) Option {
	return func(o *Options) { o.Type = s }
}
func Pass(s string) Option {
	return func(o *Options) { o.Password = s }
}
func Method(s string) Option {
	return func(o *Options) { o.Method = s }
}
func Obfs(s string) Option {
	return func(o *Options) { o.Obfs_param = s }
}
func ObfsParam(s string) Option {
	return func(o *Options) { o.Obfs_param = s }
}
func Protocol(s string) Option {
	return func(o *Options) { o.Protocol_param = s }
}
func Timeout(s int) Option {
	return func(o *Options) { o.Timeout = s }
}
func Dstaddr(s []byte) Option {
	return func(o *Options) { o.Da = s }
}

func GetOption(c Cnf) (o []Option) {
	if p := c.GetString("password"); p != "" {
		o = append(o, Pass(p))
	}
	return
}
