package home

import (
	"net/http"
	"samtake/justdoit/common"
	"samtake/justdoit/model"
	"samtake/justdoit/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

//LocalNavList  .
func LocalNavList(ctx *gin.Context) {
	icon := ctx.PostForm("icon")
	title := ctx.PostForm("title")
	url := ctx.PostForm("url")
	statusBarColor := ctx.PostForm("statusBarColor")
	hideAppBar, err := strconv.ParseBool(ctx.PostForm("hideAppBar"))
	if err != nil {
		response.Response(ctx, http.StatusServiceUnavailable, 422, nil, "参数解析错误")
	}

	if len(icon) == 0 || len(title) == 0 || len(url) == 0 || len(statusBarColor) == 0 {
		response.Response(ctx, http.StatusServiceUnavailable, 422, nil, "参数不完成")

		return
	}

	newLocalNavList := model.LocalNavList{
		Icon:           icon,
		Title:          title,
		URL:            url,
		StatusBarColor: statusBarColor,
		HideAppBar:     hideAppBar,
	}

	DB := common.GetDB()
	DB.Create(&newLocalNavList)

	//返回结果
	response.Response(ctx, http.StatusUnprocessableEntity, 200, nil, "修改成功")
}
