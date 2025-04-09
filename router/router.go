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

	//ルーティング
	mux.HandleFunc("/getusers", rt.uh.GetUserHandler)
	mux.HandleFunc("/getuser/", rt.uh.GetUserByIdHandler) // "/getuser/" パターンでマッチ
	mux.HandleFunc("/", rt.uh.HandleRoot)
	mux.HandleFunc("/signup", rt.uh.RegisterHandler)
	mux.HandleFunc("/login", rt.uh.LoginHandler)

	fmt.Println("Server Running")

	// サーバーの設定
	port := "8080"
	server := &http.Server{Addr: ":" + port, Handler: mux}

	return server
}
