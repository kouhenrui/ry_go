package route

import (
	"github.com/gin-gonic/gin"
	"ry_go/src/middleware"
)

/**
 * @ClassName router
 * @Description TODO
 * @Author khr
 * @Date 2023/7/28 15:42
 * @Version 1.0
 */
func InitRoute() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Cors())
	//r.Use(middleware.GolbalMiddleWare())          //token验证
	r.Use(middleware.LoggerMiddleWare())          //日志捕捉
	r.Use(middleware.UnifiedResponseMiddleware()) //全局统一返回格式
	//r.Use(middleware.CasbinMiddleWare())             //casbin挂载验证
	r.Use(middleware.GlobalErrorMiddleware())              //错误捕捉
	r.Use(middleware.NotFoundAndMethodNotAllowedHandler()) //404，405

	//r.R.ParseMultipartForm
	InitApi(r) //挂载请求路径
	return r
}
