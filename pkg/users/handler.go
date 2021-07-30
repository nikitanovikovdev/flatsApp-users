package users

import (
	"context"
)

type Handler struct {
	s *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		s: s,
	}
}

func (h *Handler) SignUp(ctx context.Context, username, password string) (interface{}, error) {
	id, err := h.s.CreateUser(ctx, username, password)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (h *Handler) SignIn(ctx context.Context, username, password string)  (string, error){
	token, err := h.s.GenerateToken(ctx, username, password)
	if err != nil {
		return "", err
	}

	return token, nil
}