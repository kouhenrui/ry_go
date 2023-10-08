package admin

import (
	"errors"
	"ry_go/src/dto/comDto"
	"ry_go/src/dto/reqDto"
	"ry_go/src/dto/resDto"
	"ry_go/src/global"
	"ry_go/src/pojo"
	util "ry_go/src/utils"
)

var (
	admin *pojo.Admin
	err   error
)

type AdminInter interface {
	Login(body reqDto.AdminLogin, ip string) (resDto.TokenAndExp, error)
	Register(body reqDto.AddAdmin) error
}
type AdminService struct {
	pojo.AdminRepositoryInter
}

/**
 * @Author Khr
 * @Description //TODO 登录方法
 * @Date 17:33 2023/10/5
 * @Param
 * @return
 **/
func (i *AdminService) Login(body reqDto.AdminLogin, ip string) (t resDto.TokenAndExp, err error) {
	//判定登陆方式
	switch body.Method {
	case "name":
		admin, err = i.CheckByName(body.UserName)
		break

	case "phone":
		admin, err = i.CheckByPhone(body.Phone)
		break
	default:
		return t, errors.New(util.AUTH_LOGIN_ERROR)
	}
	//将密码解密
	dePwd, err := util.DePwdCode(admin.Password, admin.Salt)
	if err != nil {
		return t, err
	}
	//解密密码和输入密码比对
	if dePwd != body.Password {
		return t, errors.New(util.ACCOUNT_PWD_ERROR)
	}
	//查询redis缓存是否存在
	redisErr := global.ExistRedis(admin.AccessToken)

	tokenDate := comDto.TokenClaims{
		Id:       admin.ID,
		Phone:    admin.Phone,
		NickName: admin.NickName,
		Role:     admin.Role,
	}
	//去除登录记录
	if body.Revoke {
		if len(admin.AccessToken) > 0 {
			if redisErr != nil {
				global.DelRedis(admin.AccessToken)
			}
			if err = i.RemoveAccessToken(admin.ID); err != nil {
				return t, err
			}
		}
		if t, err = i.tokenRedis(tokenDate, ip); err != nil {
			return t, err
		}
	} else {
		if redisErr != nil {

			if t, err = i.tokenRedis(tokenDate, ip); err != nil {
				return t, err
			}
		} else {
			tokenValue := global.GetRedis(admin.AccessToken)
			util.UnMarshal([]byte(tokenValue), &t)
		}
	}

	//t, err = i.tokenRedis(tokenDate, ip)
	//if err != nil {
	//	return t, err
	//}
	return t, nil
}

/**
 * @Author Khr
 * @Description //TODO token生成，存入缓存
 * @Date 10:43 2023/10/6
 * @Param
 * @return
 **/
func (i *AdminService) tokenRedis(tokenDate comDto.TokenClaims, ip string) (t resDto.TokenAndExp, err error) {
	t = util.SignToken(tokenDate, global.AdminExp)
	//生成access_token
	access_token := util.Rand6String()
	//redis缓存token
	if err = global.SetRedis(access_token, util.Marshal(t), global.AdminExp); err != nil {
		return t, err
	}
	//数据库保存access_token
	if err = i.UpdateToken(access_token, admin.ID, ip); err != nil {
		return t, err
	}
	return t, nil
}

/**
 * @Author Khr
 * @Description //TODO 添加数据
 * @Date 11:59 2023/10/6
 * @Param
 * @return
 **/
func (i *AdminService) Register(body reqDto.AddAdmin) error {

	return nil
}
