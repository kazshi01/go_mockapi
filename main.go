package main

import (
	"mockapi/model"
	"mockapi/repository"
	"mockapi/router"
	"mockapi/service"
	"mockapi/web"
)

func main() {
	model := &model.User{Name: "Kazuma", Age: 35}
	ur := repository.NewUserRepository(model)
	us := service.NewUserService(ur)
	uh := web.NewUserHandler(us)
	rt := router.NewUserRouter(uh)

	server := rt.SetupRoutes()
	server.ListenAndServe()
}
