package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ry_go/src/dto/reqDto"
	"ry_go/src/msg"
	"ry_go/src/service/admin"
	util "ry_go/src/utils"
)

var (
	adminService admin.AdminInter = &admin.AdminService{}
	//adminService admin.AdminService
)

func Login(c *gin.Context) {
	var adminLogin reqDto.AdminLogin
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

func LogOut(c *gin.Context) {
	var adminRegister reqDto.AddAdmin
	if err := c.ShouldBindJSON(&adminRegister); err != nil {
		c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &adminRegister))
		return
	}

}

func Info(c *gin.Context) {
	c.Set("res", "")
}
func Register(c *gin.Context) {
	var addAdmin reqDto.AddAdmin
	if err := c.ShouldBindJSON(&addAdmin); err != nil {
		c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &addAdmin))
		return
	}
	err := adminService.Register(addAdmin)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("res", msg.ADD_SUCCESS)

}
