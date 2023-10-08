package route

import (
	"github.com/gin-gonic/gin"
	"ry_go/src/controller/common"
)

/**
 * @ClassName api
 * @Description TODO
 * @Author khr
 * @Date 2023/7/29 14:18
 * @Version 1.0
 */

func InitApi(route *gin.Engine) {
	api := route.Group("/api")
	{

		api.GET("/captcha", common.GetCaptcha)
		api.POST("/verify/captcha", common.VfCaptcha)
		api.POST("/upload/file", common.UploadFile)
		//authModule := api.Group("/auth")
		//
		//{
		//	authModule.POST("auth")
		//}
	}

}
