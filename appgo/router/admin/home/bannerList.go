package home

import (
	"net/http"
	"samtake/justdoit/common"
	"samtake/justdoit/model"
	"samtake/justdoit/response"

	"github.com/gin-gonic/gin"
)

//BannerList 首页banner.
func BannerList(ctx *gin.Context) {
	icon := ctx.PostForm("icon")
	url := ctx.PostForm("url")

	if len(icon) == 0 {
		response.Response(ctx, http.StatusServiceUnavailable, 422, nil, "icon不能为空")

		return
	}

	if len(url) == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "url不能为空")
		return
	}

	newBannerList := model.BannerList{
		Icon: icon,
		URL:  url,
	}

	DB := common.GetDB()
	DB.Create(&newBannerList)

	//返回结果
	response.Success(ctx, nil, "修改成功")
}
