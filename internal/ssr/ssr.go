package ssr

import (
	"github.com/jWhisper/ssrlocal/configs"
	"github.com/jWhisper/ssrlocal/errs"
	"github.com/jWhisper/ssrlocal/internal/ssr/obfs"
	"github.com/jWhisper/ssrlocal/internal/ssr/pro"
	"github.com/jWhisper/ssrlocal/pkg/log"
)

var (
	logger       = log.WithLevelAndMeta(log.DefaultLogger, log.LvInfo, "ssrlocal:")
	tcpKeepAlive = false
	tcpSndBuf    = 4 * 1024
	tcpRcvBuf    = 4 * 1024
)

type server struct {
	addrs     []string
	t, rp, lp string
	cnf       configs.Cnf
	obfs      obfs.Obfs
	pro       pro.Protocol
	err       error
}

func NewServer(cnf configs.Cnf) (s *server, err error) {
	t := cnf.GetString("type")
	lp, ok := cnf.Get("local_port").(string)
	if !ok {
		return nil, errs.InvalidLocalPort
	}
	rp, ok := cnf.Get("server_port").(string)
	if !ok {
		return nil, errs.InvalidRemotePort
	}
	addrs := cnf.GetStringSlice("server")
	if len(addrs) < 1 {
		return nil, errs.InvalidRemoteServers
	}
	logger.Print("listen at", lp, ";", "remote servers:", addrs, "remote port", rp)

	s = &server{
		t:     t,
		cnf:   cnf,
		addrs: addrs, // a server slice
		rp:    rp,
		lp:    lp,
		obfs:  obfs.ObfsImp{},
		pro:   pro.ProImp{},
		err:   err,
	}
	return
}
