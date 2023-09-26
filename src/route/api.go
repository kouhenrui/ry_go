package route

import (
	"github.com/gin-gonic/gin"
	"ry_go/src/controller/commonController"
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

		api.GET("/captcha", commonController.GetCaptcha)
		api.POST("/verify/captcha", commonController.VfCaptcha)
		//authModule := api.Group("/auth")
		//
		//{
		//	authModule.POST("auth")
		//}
	}

}
