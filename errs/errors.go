package errs

import "errors"

var (
	InvalidLocalPort       = errors.New("invalid local port to listen")
	InvalidRemotePort      = errors.New("invalid remote port to listen")
	InvalidRemoteServers   = errors.New("invalid remote servers")
	ErrAddressNotSupported = errors.New("invalid dst address")
	ErrCommandNotSupported = errors.New("invalid command")

	ErrCipherMethodNotSupported = errors.New("invalid cipher method")
	ErrNilConn                  = errors.New("nil conn")

	ErrObfsNotSupported     = errors.New("invalid obfs")
	ErrProtocolNotSupported = errors.New("invalid protocol")
)
