package util

import (
	"github.com/gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"strconv"
)


func GetPage(c *gin.Context) int {
	// 获取分页页码
	result := 0
	//page, _ := com.StrTo(c.Query("page")).Int()
	page,_ := strconv.Atoi(c.Query("page")) // 将str转换成int

	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
