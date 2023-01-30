package postgre

import (
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler(host string, port int, username, password string) (*Handler, error) {
	return &Handler{}, nil
}

func (h *Handler) Handle(subRouter *mux.Router) {
}
