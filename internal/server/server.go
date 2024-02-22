package server

import "github.com/NotFound1911/mserver"

type Server struct {
	Addr string
	*mserver.Core
}
