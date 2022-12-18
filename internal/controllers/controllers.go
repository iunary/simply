package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/iunary/simply/internal/transports/http"
)

var ProviderSet = wire.NewSet(NewUserController, SetupHandlers)

func SetupHandlers(uc *UserController) http.Handlers {
	return func(r *gin.Engine) {
		users := r.Group("/users")
		users.GET("", uc.Get)
	}
}
