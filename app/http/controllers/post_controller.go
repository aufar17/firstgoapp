package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
  	"goravel/app/models"

)

type PostController struct {
	// Dependent services
}

func NewPostController() *PostController {
	return &PostController{
		// Inject services
	}
}

func (r *PostController) Index(ctx http.Context) http.Response {
	var posts []models.Post
	facades.Orm().Query().Find(&posts)
	return ctx.Response().Success().Json(http.Json{"posts": posts})	


}	
