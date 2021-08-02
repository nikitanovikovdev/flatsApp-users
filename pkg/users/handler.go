package users

import (
	"encoding/json"
	"github.com/nikitanovikovdev/flatsApp-users/pkg/platform/response"
	"github.com/nikitanovikovdev/flatsApp-users/pkg/platform/user"
	"net/http"
)



type Handler struct {
	s *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		s: s,
	}
}

func (h *Handler) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user user.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			response.UserError(w, err)
			return
		}

		id, err := h.s.CreateUser(r.Context(), user)
		if err != nil {
			response.UserError(w, err)
			return
		}

		idStr := id.(string)

		response.OkWithMessage(w, []byte(idStr))
	}
}

func (h *Handler) SignIn()  http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var user user.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			response.UserError(w, err)
			return
		}

		token, err := h.s.GenerateToken(r.Context(), user)
		if err != nil {
			response.DevError(w, err)
			return
		}

		response.OkWithMessage(w, []byte(token))
	}
}


