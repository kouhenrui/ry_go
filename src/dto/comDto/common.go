package comDto

import "ry_go/src/pojo"

/**
 * @ClassName common
 * @Description TODO
 * @Author khr
 * @Date 2023/5/8 13:37
 * @Version 1.0
 */

type TokenClaims struct {
	Id       uint        `json:"id"`
	Phone    string      `json:"phone"`
	NickName string      `json:"nick_name"`
	Role     []pojo.Role `json:"role"`
}
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
