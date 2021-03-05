package errs

import "errors"

var (
	InvalidLocalPort     = errors.New("invalid local port to listen")
	InvalidRemotePort    = errors.New("invalid remote port to listen")
	InvalidRemoteServers = errors.New("invalid remote servers")
)
