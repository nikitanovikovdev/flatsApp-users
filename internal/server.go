package internal

import (
	"context"
	"fmt"
	"github.com/nikitanovikovdev/flatsApp-users/pkg/users"
	auth "github.com/nikitanovikovdev/flatsApp-users/proto"
)

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

func (g *GRPCServer) Registr(ctx context.Context, req *auth.RegistrData) (*auth.Id, error) {
	idRes, err := g.h.SignUp(ctx, req.Username, req.Password)

	id, ok := idRes.(string)
	if !ok {
		return nil, err
	}

	if err != nil {
		fmt.Sprintf("invalid user :%v", err)
	}
	return &auth.Id{Id: id}, nil
}
