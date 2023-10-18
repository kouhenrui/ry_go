package admin

import (
	"errors"
	"fmt"
	"log"
	"ry_go/src/dto/comDto"
	"ry_go/src/dto/reqDto"
	"ry_go/src/dto/resDto"
	"ry_go/src/global"
	inter2 "ry_go/src/inter"
	"ry_go/src/msg"
	"ry_go/src/pojo"
	"ry_go/src/service/common"
	util "ry_go/src/utils"
)

var (
	admin           *pojo.Admin
	err             error
	commonService   common.CommonServiceImpl
	adminRepository inter2.AdminRepositoryInter = &inter2.AdminRepositoryImpl{}
	roleRepository  inter2.RoleInter            = &inter2.RoleRepositoryImpl{}
)

type AdminInter interface {
	Login(body reqDto.AdminLogin, ip string) (resDto.TokenAndExp, error)
	Register(body reqDto.AddAdmin) error
	Info(id uint) (*resDto.AdminInfo, error)
}
type AdminService struct{}

// TODO 登录方法
func (i *AdminService) Login(body reqDto.AdminLogin, ip string) (t resDto.TokenAndExp, err error) {
	//captcha := &reqDto.Captcha{
	//	Id:      body.Uuid,
	//	Content: body.Code,
	//}
	//if err = global.ExistRedis(captcha.Id); err == nil {
	//	return t, errors.New(msg.CAPTCHA_ERROR)
	//}
	//if !commonService.VfCaptcha(*captcha) {
	//	return t, errors.New(msg.CAPTCHA_CREATE_ERROR)
	//}
	//判定登陆方式
	switch body.Method {
	case "name":
		admin, err = adminRepository.CheckByName(body.UserName)
		break

	case "phone":
		admin, err = adminRepository.CheckByPhone(body.Phone)
		break
	default:
		return t, errors.New(msg.AUTH_LOGIN_ERROR)
	}
	//将密码解密
	dePwd, err := util.DePwdCode(admin.Password, admin.Salt)
	if err != nil {
		return t, err
	}
	//解密密码和输入密码比对
	if dePwd != body.Password {
		return t, errors.New(msg.ACCOUNT_PWD_ERROR)
	}
	//查询redis缓存是否存在
	redisErr := global.ExistRedis(admin.AccessToken)
	tokenDate := comDto.TokenClaims{
		Id:       admin.ID,
		Phone:    admin.Phone,
		Name:     admin.UserName,
		NickName: admin.NickName,
		Role:     admin.Roles,
	}
	//去除登录记录
	if body.Revoke {
		if len(admin.AccessToken) > 0 {
			if redisErr == nil {
				_ = global.DelRedis(admin.AccessToken)
			}
			err = adminRepository.RemoveAccessToken(admin.ID)
		}
		t, err = i.tokenRedis(tokenDate, ip)
	} else {
		if len(admin.AccessToken) > 0 {
			if redisErr != nil {
				t, err = i.tokenRedis(tokenDate, ip)
			} else {
				tokenValue := global.GetRedis(admin.AccessToken)
				util.UnMarshal([]byte(tokenValue), &t)
			}
		} else {
			t, err = i.tokenRedis(tokenDate, ip)
		}
	}

	return t, nil
}

// TODO info
func (i *AdminService) Info(id uint) (*resDto.AdminInfo, error) {
	var adminInfo *pojo.Admin
	adminInfo, err = adminRepository.AdminInfo(id)
	var info = &resDto.AdminInfo{
		Id:       adminInfo.ID,
		UserName: adminInfo.UserName,
		NickName: adminInfo.NickName,
		Phone:    adminInfo.Phone,
		Sex:      adminInfo.Sex,
		Avatar:   adminInfo.Avatar,
		Email:    adminInfo.Email}
	for _, i := range adminInfo.Roles {
		r := i.Name
		info.Role = append(info.Role, r)
	}
	return info, err
}

// TODO token生成，存入缓存
func (i *AdminService) tokenRedis(tokenDate comDto.TokenClaims, ip string) (t resDto.TokenAndExp, err error) {
	t = util.SignToken(tokenDate, global.AdminExp)
	//生成access_token
	access_token := util.Rand6String()
	//redis缓存token
	if err = global.SetRedis(access_token, util.Marshal(t), global.AdminExp); err != nil {
		return t, err
	}
	//数据库保存access_token
	if err = adminRepository.UpdateToken(access_token, admin.ID, ip); err != nil {
		return t, err
	}
	return t, nil
}

// TODO 添加数据
func (i *AdminService) Register(body reqDto.AddAdmin) error {
	if len(body.UserName) < 0 && len(body.Phone) < 0 {
		return errors.New(msg.ACCOUNT_PHONE_NOT_NULL)
	}
	if len(body.UserName) > 0 {
		fmt.Println(body.UserName)
		_, err = adminRepository.CheckByName(body.UserName)
		if err != nil {
			return err
		}
	}
	if len(body.Phone) > 0 {
		_, err = adminRepository.CheckByPhone(body.Phone)
		if err != nil {
			return err
		}
	}
	var salt = util.MakeSalt()
	var enpwd, _ = util.EnPwdCode(body.Password, salt)
	var roles = []pojo.Role{}
	if len(body.Role) > 0 {
		for _, id := range body.Role {
			var singleRole pojo.Role
			findRole, signErr := roleRepository.FindById(id)
			if signErr != nil {
				return err
			}
			singleRole.Model = findRole.Model
			singleRole.Name = findRole.Name
			roles = append(roles, singleRole)

		}
	}
	var newAdmin = pojo.Admin{
		UserName: body.UserName,
		NickName: body.NickName,
		Password: enpwd,
		Salt:     salt,
		Phone:    body.Phone,
		Sex:      body.Sex,
		Avatar:   body.Avatar,
		Email:    body.Email,
		Roles:    roles,
	}
	log.Println("打印要插入的数据", newAdmin)
	return adminRepository.AddAdmin(newAdmin)
}
