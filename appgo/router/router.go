package router

import (
	"samtake/justdoit/controller"

	"github.com/gin-gonic/gin"

	"samtake/justdoit/middleware"
)

//CollectRoute .
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
