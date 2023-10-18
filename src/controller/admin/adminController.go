package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ry_go/src/dto/reqDto"
	"ry_go/src/dto/resDto"
	"ry_go/src/msg"
	"ry_go/src/service/admin"
	util "ry_go/src/utils"
	"time"
)

var (
	adminService admin.AdminInter = &admin.AdminService{}
	err          error
)

func Login(c *gin.Context) {
	var adminLogin reqDto.AdminLogin
	if err = c.ShouldBindJSON(&adminLogin); err != nil {
		c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &adminLogin))
		return
	}
	ip := c.GetString("ip")
	t, err := adminService.Login(adminLogin, ip)
	if err != nil {
		c.Error(err)
	}
	// 将Token设置到Cookie中
	c.SetCookie("token", t.Token, int(time.Hour*24*7), c.GetString("reqUrl"), ip, false, true)
	c.Set("res", t)
}

func LogOut(c *gin.Context) {

	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.Set("res", "ok")
	//var adminRegister reqDto.AddAdmin
	//if err = c.ShouldBindJSON(&adminRegister); err != nil {
	//	c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &adminRegister))
	//	return
	//}

}

func Info(c *gin.Context) {
	id := c.GetUint("user_id")
	fmt.Println(id, "id")
	var info = &resDto.AdminInfo{}
	info, err = adminService.Info(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("res", info)
}
func Register(c *gin.Context) {
	var addAdmin reqDto.AddAdmin
	if err = c.ShouldBindJSON(&addAdmin); err != nil {
		c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &addAdmin))
		return
	}
	err = adminService.Register(addAdmin)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("res", msg.ADD_SUCCESS)

}
