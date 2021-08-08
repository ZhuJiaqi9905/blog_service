package app

import (
	"blog-service/global"
	"blog-service/pkg/convert"

	"github.com/gin-gonic/gin"
)

func GetPage(ctx *gin.Context) int {
	page := convert.StrTo(ctx.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(ctx *gin.Context) int {
	pageSize := convert.StrTo(ctx.Query(("page_size"))).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := (page - 1) * pageSize
	return result
}
