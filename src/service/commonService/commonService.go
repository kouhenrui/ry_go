package commonService

import (
	"ry_go/src/dto/reqDto"
	util "ry_go/src/utils"
)

var (
	err    error
	capcha = util.NewCaptchaService()
)

func GetCaptcha() (error, reqDto.Captcha) {

	var newCaptcha reqDto.Captcha

	newCaptcha.Id, newCaptcha.Content, err = capcha.GenerateCaptcha()
	if err != nil {
		return err, newCaptcha
	}
	return nil, newCaptcha
}
func VfCaptcha(capt reqDto.Captcha) bool {
	return capcha.VerifyCaptcha(capt.Id, capt.Content)
}
