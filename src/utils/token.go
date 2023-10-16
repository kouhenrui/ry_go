package util

import (
	"fmt"
	"github.com/go-pay/gopay/pkg/jwt"
	"ry_go/src/dto/comDto"
	"ry_go/src/dto/resDto"
	"ry_go/src/global"
	"time"
)

var jwtkey = []byte(global.JWTKEY)

// var userInfo pojo.User
type UserClaims struct {
	Name     string `json:"name"`
	Role     int    `json:"role"`
	Account  string `json:"account"`
	Id       uint   `json:"id"`
	RoleName string `json:"role_name"`
}

type AllClaims struct {
	jwt.StandardClaims
	User comDto.TokenClaims
}

// 颁发token inter
func SignToken(infoClaims comDto.TokenClaims, day time.Duration) (t resDto.TokenAndExp) {
	expireTime := time.Now().Add(day) //7天过期时间
	claims := &AllClaims{
		User: infoClaims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "khr",  // 签名颁发者
			Subject:   "sign", //签名主题
		},
	}
	//fmt.Println(claims, "封装的信息")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err, "生成token错误")
	}
	tFStr := expireTime.Format("2006-01-02 15:04:05")
	t.Token = tokenString
	t.Exptime = tFStr
	return t
}

//// 验证token
//func AnalysyToken(c *gin.Context) bool {
//	//fmt.Println("进入token验证")
//	tokenString := c.GetHeader("Authorization")
//	if tokenString == "" {
//		return false
//	}
//	return true
//}

// 解析Token
func ParseToken(tokenString string) comDto.TokenClaims {
	//claims := &Claims{}
	//解析token
	token, _ := jwt.ParseWithClaims(tokenString, &AllClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	user, _ := token.Claims.(*AllClaims)
	//fmt.Println(user.User, "打印")
	return user.User
}
