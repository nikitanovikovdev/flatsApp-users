package internal

import (
	"context"
	"fmt"
	"github.com/nikitanovikovdev/flatsApp-users/pkg/users"
	auth "github.com/nikitanovikovdev/flatsApp-users/proto"
)

//type Server struct {
//	srv http.Server
//}
//
//func NewServer(host, port string, h http.Handler) *Server {
//	return &Server{
//		srv: http.Server{
//			Addr:    host + ":" + port,
//			Handler: h,
//		},
//	}
//}
//
//func (s *Server) Run() error {
//	return s.srv.ListenAndServe()
//}
//
//func(s *Server) Shutdown(ctx context.Context) error {
//	return s.srv.Shutdown(ctx)
//}


type GRPCServer struct {
	h *users.Handler
}

func NewGRPCServer(h *users.Handler) *GRPCServer{
	return &GRPCServer{
		h: h,
	}
}

func (g *GRPCServer) Authorize(ctx context.Context, req *auth.RequestData) (*auth.Token, error) {
	token, err := g.h.SignIn(ctx, req.Username, req.Password)
	if err != nil {
		fmt.Sprintf("invalid user :%v", err)
	}
	return &auth.Token{Token: token}, nil
}