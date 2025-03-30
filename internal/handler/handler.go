package handler

import "github.com/sklevenz/cf-api-server/internal/generated"

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ generated.ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}
