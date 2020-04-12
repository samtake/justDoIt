package home

import (
	"net/http"
	"samtake/justdoit/common"
	"samtake/justdoit/model"
	"samtake/justdoit/response"

	"github.com/gin-gonic/gin"
)

//SalesBox .
func SalesBox(ctx *gin.Context) {
	icon := ctx.PostForm("icon")
	url := ctx.PostForm("url")
	title := ctx.PostForm("title")
	typeName := ctx.PostForm("type")

	if len(icon) == 0 || len(url) == 0 || len(title) == 0 || len(typeName) == 0 {
		response.Response(ctx, http.StatusServiceUnavailable, 422, nil, "参数不齐")

		return
	}

	newSalesBox := model.SalesBox{
		Icon:     icon,
		Title:    title,
		URL:      url,
		TypeName: typeName,
	}

	common.GetDB().Create(&newSalesBox)

	//返回结果
	response.Response(ctx, http.StatusUnprocessableEntity, 200, nil, "修改成功")
}
