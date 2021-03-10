package obfs

type Setor func(*obfs)

func SetHost(s string) Setor {
	return func(o *obfs) { o.host = s }
}
func SetPort(s uint16) Setor {
	return func(o *obfs) { o.port = s }
}

func SetParam(s string) Setor {
	return func(o *obfs) { o.param = s }
}
func SetIV(s []byte) Setor {
	return func(o *obfs) { o.iV = s }
}
func SetIVLen(s int) Setor {
	return func(o *obfs) { o.iVLen = s }
}
func SetrecvIV(s []byte) Setor {
	return func(o *obfs) { o.recvIV = s }
}
func SetrecvIVLen(s int) Setor {
	return func(o *obfs) { o.recvIVLen = s }
}
func SetKey(s []byte) Setor {
	return func(o *obfs) { o.key = s }
}
func SetKeyLen(s int) Setor {
	return func(o *obfs) { o.keyLen = s }
}
func SetHeadLen(s int) Setor {
	return func(o *obfs) { o.headLen = s }
}
func SetTcpMss(s int) Setor {
	return func(o *obfs) { o.tcpMss = s }
}
