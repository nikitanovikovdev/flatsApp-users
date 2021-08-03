package internal

import (
	"context"
	"github.com/nikitanovikovdev/flatsApp-users/pkg/users"
	authorization "github.com/nikitanovikovdev/flatsApp-users/proto"
	"github.com/spf13/viper"
	"net/http"
)

type Server struct {
	srv http.Server
}

func NewServer(host, port string, h http.Handler) *Server {
	return &Server{
		srv: http.Server{
			Addr: host + ":" + port,
			Handler: h,
		},
	}
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

type GRPCServer struct {
	h *users.Handler
}

func (g *GRPCServer) ReturnSignKey(ctx context.Context, empty *authorization.Empty) (*authorization.SigningKey, error) {
	return &authorization.SigningKey{SigningKey: viper.GetString("keys.signing_key")}, nil
}