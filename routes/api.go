package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)

	PostController := controllers.NewPostController()
	facades.Route().Get("posts", PostController.Index)
	facades.Route().Post("/post/create", PostController.Create)
	facades.Route().Post("/post/update", PostController.Update)
	facades.Route().Post("/post/delete", PostController.Delete)

}
