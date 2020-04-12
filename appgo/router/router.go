package router

import (
	"github.com/gin-gonic/gin"

	"samtake/justdoit/middleware"
	"samtake/justdoit/router/admin/home"
	"samtake/justdoit/router/api/auth"
)

// //这里想着：通过接口来管理路由的接口方法
// //先注释，后续再做修改
// type HandlerFunc interface {}

//InitRouter .
func InitRouter(r *gin.Engine) *gin.Engine {
	webGrp(r)
	apiGrp(r)
	adminGrp(r)
	return r
}

//webGrp 小程序api路由组 .
func webGrp(r *gin.Engine) {
}

//apiGrp app api路由组 .
func apiGrp(r *gin.Engine) {
	r.POST("/api/auth/register", auth.Register)
	r.POST("/api/auth/login", auth.Login)
	r.POST("/api/auth/info", middleware.AuthMiddleware(), auth.Info)
}

//adminGrp 后台管理 api路由组.
func adminGrp(r *gin.Engine) {
	r.POST("/admin/home/bannerList", home.BannerList)
	r.POST("/admin/home/gridNavMain", home.GridNavMain)
	r.POST("/admin/home/gridNavItem", home.GridNavItem)
	r.POST("/admin/home/localNavList", home.LocalNavList)
	r.POST("/admin/home/salesBox", home.SalesBox)
	r.POST("/admin/home/subNavList", home.SubNavList)
}
