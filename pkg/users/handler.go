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

//func (h *Handler) SignUp() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		body, err := ioutil.ReadAll(r.Body)
//		if err != nil {
//			response.UserError(w, err)
//			return
//		}
//
//		id, err := h.s.CreateUser(r.Context(), body)
//		if err != nil {
//			response.DevError(w, err)
//			return
//		}
//
//		message, err := json.Marshal(id)
//		if err != nil {
//			response.DevError(w, err)
//			return
//		}
//
//		response.CreateWithMessage(w, message)
//	}
//}

func (h *Handler) SignIn(ctx context.Context, username, password string)  (string, error){
	token, err := h.s.GenerateToken(ctx, username, password)
	if err != nil {
		return "", err
	}

	return token, nil
}