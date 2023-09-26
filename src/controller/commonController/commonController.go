package commonController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ry_go/src/dto/reqDto"
	"ry_go/src/service/commonService"
	util "ry_go/src/utils"
)

func GetCaptcha(c *gin.Context) {
	err, captera := commonService.GetCaptcha()
	//fmt.Println("接获获取", captera)
	if err != nil {
		c.Error(err)
		return
	}

	c.Set("res", captera)

}

func VfCaptcha(c *gin.Context) {
	var vf = reqDto.Captcha{}
	if err := c.ShouldBindJSON(&vf); err != nil {
		c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &vf))
		return
	}
	c.Set("res", commonService.VfCaptcha(vf))
	//return commonService.VfCaptcha(vf)
}
