package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ry_go/src/dto/reqDto"
	"ry_go/src/service/admin"
	util "ry_go/src/utils"
)

var (
	adminService admin.AdminInter
)

func Login(c *gin.Context) {
	var adminLogin = reqDto.AdminLogin{}
	if err := c.ShouldBindJSON(&adminLogin); err != nil {
		c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &adminLogin))
		return
	}
	ip := c.GetString("ip")
	t, err := adminService.Login(adminLogin, ip)
	if err != nil {
		c.Error(err)
	}
	c.Set("res", t)
}

func Info(c *gin.Context) {
	c.Set("res", "")
}
func Register(c *gin.Context) {

}
