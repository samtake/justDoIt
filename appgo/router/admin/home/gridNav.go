package home

import (
	"net/http"
	"samtake/justdoit/common"
	"samtake/justdoit/model"
	"samtake/justdoit/response"

	"github.com/gin-gonic/gin"
)

//GridNavMain  .
func GridNavMain(ctx *gin.Context) {
	startColor := ctx.PostForm("startColor")
	endColor := ctx.PostForm("endColor")
	typeName := ctx.PostForm("type")

	if len(startColor) == 0 || len(endColor) == 0 || len(typeName) == 0 {
		response.Response(ctx, http.StatusServiceUnavailable, 422, nil, "参数不齐")

		return
	}

	newGridNavSub := model.GridNavSub{
		StartColor: startColor,
		EndColor:   endColor,
		TypeName:   typeName,
	}

	common.GetDB().Create(&newGridNavSub)

	//返回结果
	response.Success(ctx, nil, "修改成功")
}

//GridNavItem  .
func GridNavItem(ctx *gin.Context) {
	icon := ctx.PostForm("icon")
	title := ctx.PostForm("title")
	url := ctx.PostForm("url")
	statusBarColor := ctx.PostForm("statusBarColor")
	typeName := ctx.PostForm("type")

	if len(icon) == 0 || len(title) == 0 || len(url) == 0 || len(statusBarColor) == 0 || len(typeName) == 0 {
		response.Response(ctx, http.StatusServiceUnavailable, 422, nil, "参数不齐")

		return
	}

	newGridNavSubMainItem := model.GridNavSubMainItem{
		Icon:           icon,
		Title:          title,
		URL:            url,
		StatusBarColor: statusBarColor,
		TypeName:       typeName,
	}

	common.GetDB().Create(&newGridNavSubMainItem)

	//返回结果
	response.Success(ctx, nil, "修改成功")
}
