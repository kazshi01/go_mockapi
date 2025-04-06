package main

import (
	"mockapi/repository"
	"mockapi/router"
	"mockapi/service"
	"mockapi/web"
)

func main() {
	ur := repository.NewUserRepository()
	us := service.NewUserService(ur)
	uh := web.NewUserHandler(us)
	rt := router.NewUserRouter(uh)

	server := rt.SetupRoutes()
	server.ListenAndServe()
}
