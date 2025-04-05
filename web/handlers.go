package web

import (
	"mockapi/service"
	"net/http"
)

type IUserHandler interface {
	SayHandler(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	us service.IUseServeice
}

func NewUserHandler(us service.IUseServeice) *UserHandler {
	return &UserHandler{us}
}

func (uh *UserHandler) SayHandler(w http.ResponseWriter, r *http.Request) {
	message := uh.us.SayPlus()
	w.Write([]byte(message))
}
