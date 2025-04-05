package router

import (
	"fmt"
	"mockapi/web"
	"net/http"
)

type UserRouter struct {
	uh web.IUserHandler
}

func NewUserRouter(uh web.IUserHandler) *UserRouter {
	return &UserRouter{uh}
}

func (rt *UserRouter) SetupRoutes() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/say", rt.uh.SayHandler)
	fmt.Println("Server Running")

	// サーバーの設定
	port := "8080"
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	return server
}
