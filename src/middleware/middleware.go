package middleware

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"io/ioutil"
	"math"
	"net/http"
	"ry_go/src/dto/comDto"
	"ry_go/src/global"
	util "ry_go/src/utils"
	"sync"
	"time"
)

/**
 * @ClassName middleware
 * @Description 中间件
 * @Author khr
 * @Date 2023/7/28 15:45
 * @Version 1.0
 */

var (
	requestCounts = make(map[string]int)
	claims        comDto.TokenClaims
)

/*
* @MethodName Cors
* @Description 跨域，限制请求方法，限制请求头
* @Author khr
* @Date 2023/7/29 9:52
 */

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Origin, X-CSRF-Token,X-Requested-With,Accept, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// 允许放行OPTIONS请求
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func NotFoundAndMethodNotAllowedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 如果在处理请求后仍然没有找到路由，说明是404错误
		if c.Writer.Status() == http.StatusNotFound {
			c.AbortWithError(http.StatusNotFound, errors.New("404 Not Found"))
		}

		// 如果请求方法不被允许，说明是405错误
		if c.Writer.Status() == http.StatusMethodNotAllowed {
			c.AbortWithError(http.StatusMethodNotAllowed, errors.New("405 Method Not Allowed"))
		}
	}
}

//func MethodNotAllowedHandler() gin.HandlerFunc {
//	fmt.Println("405不允许")
//	return func(c *gin.Context) {
//		c.AbortWithError(http.StatusMethodNotAllowed, errors.New("405 Method Not Allowed"))
//	}
//}
//
//func NotFoundHandler() gin.HandlerFunc {
//	fmt.Println("404未找到")
//	//c.AbortWithError(http.StatusNotFound, errors.New("404 Not Found"))
//	return func(c *gin.Context) {
//		c.AbortWithError(http.StatusNotFound, errors.New("404 Not Found"))
//	}
//}

/*
* @MethodName
* @Description 日志中间件
* @Author khr
* @Date 2023/7/31 15:19
 */

func LoggerMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		requestBody, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		rbody := string(requestBody)
		query := c.Request.URL.RawQuery
		c.Next() // 调用该请求的剩余处理程序
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000))))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		url := c.Request.RequestURI
		Log := global.Logger.WithFields(
			logrus.Fields{
				"SpendTime": spendTime,       //接口花费时间
				"path":      url,             //请求路径
				"Method":    method,          //请求方法
				"status":    statusCode,      //接口返回状态
				"proto":     c.Request.Proto, //http请求版本
				"Ip":        clientIP,        //IP地址
				"body":      rbody,           //请求体
				"query":     query,           //请求query
				"message":   c.Errors,        //返回错误信息
			})
		if len(c.Errors) > 0 { // 矿建内部错误
			Log.Error(c.Errors.ByType(gin.ErrorTypePrivate))
		}
		if statusCode > 200 {
			Log.Error()
		} else {
			Log.Info()
		}
	}
}

/*
* @MethodName
* @Description token验证
* @Author khr
* @Date 2023/7/31 16:37
 */
func GolbalMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("token认证开始执行")
		//t := time.Now()
		requestUrl := c.Request.URL.String()
		//路径模糊匹配
		if !util.FuzzyMatch(requestUrl, global.WriteList) {
			//请求头是否携带token
			judge := util.AnalysyToken(c)
			if !judge {
				c.AbortWithStatusJSON(http.StatusUnauthorized, util.NO_AUTHORIZATION)
				return
			}
			claims = util.ParseToken(c.GetHeader("Authorization"))
			c.Set("role_name", claims)
		}
		c.Next()
		//ts := time.Since(t)
		//fmt.Println("time", ts)
		//fmt.Println("token认证执行结束")

	}
}

/*
 * @MethodName
 * @Description 权限路由验证
 * @Author khr
 * @Date 2023/7/31 16:35
 */

//func CasbinMiddleWare() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		requestUrl := c.Request.URL.String()
//		reqUrl := strings.Split(requestUrl, "/api/")
//		rolename := global.RoleName
//		paths := global.ReuqestPaths
//		pathIsExist := util.ExistIn(reqUrl[1], paths)
//		//登录跳过权限验证
//		if !pathIsExist {
//			//验证身份
//			_, y := c.Get("ok")
//			//通过身份验证
//			if !y {
//				c.AbortWithStatusJSON(http.StatusUnauthorized, util.NO_AUTH_ERROR)
//				return
//			} else {
//				roleName := c.GetString("role_name")
//				role := c.GetInt("role")
//				if !util.ExistIn(roleName, rolename) {
//					err, permission := permissionServiceImpl.FindPermissionByPath(reqUrl[1])
//					if err != nil {
//						c.AbortWithStatusJSON(http.StatusAccepted, util.INSUFFICIENT_PERMISSION_ERROR)
//						return
//					}
//					allowRole := permission.AuthorizedRoles
//					roleList := strings.Split(allowRole, ",")
//					roleExist := util.ExistIn(string(role), roleList)
//					if !roleExist {
//						//c.Abort()
//						//fmt.Println("请求地址不包含该权限权限")
//						c.AbortWithStatusJSON(http.StatusAccepted, util.INSUFFICENT_PERMISSION)
//						//res.Err(util.INSUFFICENT_PERMISSION)
//						return
//					}
//				}
//				fmt.Println("检测到是超级管理员，可以直接操作，不需要判断")
//			}
//		}
//	}
//}
//func MethodNotAllowedHandler(c *gin.Context) {
//	fmt.Println("405不允许")
//	c.AbortWithError(http.StatusMethodNotAllowed, errors.New("405 Method Not Allowed"))
//	return
//}
//func NotFoundHandler(c *gin.Context) {
//	fmt.Println("404未找到")
//	c.AbortWithError(http.StatusNotFound, errors.New("404 Not Found"))
//	return
//}

//type ValidationErrors map[string]string
//
//func NewValidationErrors(err validator.ValidationErrors) ValidationErrors {
//	validationErrors := make(ValidationErrors)
//	for _, e := range err {
//		validationErrors[e.Field()] = e.Tag()
//	}
//	return validationErrors
//}
//
//func validateParams(c *gin.Context) {
//	if err := c.ShouldBindQuery(c); err != nil {
//		if _, ok := err.(validator.ValidationErrors); ok {
//			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "errors": NewValidationErrors(err.(validator.ValidationErrors))})
//			return
//		} else {
//			c.JSON(http.StatusInternalServerError, gin.H{
//
//				"error": "Internal Server Error"})
//			c.Abort()
//			return
//		}
//	}
//}

/*
 * @MethodName
 * @Description IP请求次数拦截中间件,拦截ip
 * @Author khr
 * @Date 2023/7/31 15:24
 */
var visitorMap = make(map[string]*rate.Limiter) // 存储IP地址和速率限制器的映射
var mu sync.Mutex                               // 互斥锁，保证并发安全

//func IPInterceptor() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		ip := c.ClientIP()
//		if ip == "" {
//			ip = c.Request.RemoteAddr
//		}
//		if util.ExistIn(ip, global.IpAccess) {
//			c.Next()
//		}
//		path := c.Request.URL.Path
//		//fmt.Println(ip, path)
//
//		// 组合出 key
//		key := fmt.Sprintf("request:%s:%s", ip, path)
//		//fmt.Print("key", key)
//		// 将请求次数 +1，并设置过期时间
//		err := global.AutoInc(key)
//
//		if err != nil {
//			// 记录日志
//			fmt.Println("incr error:", err)
//			c.AbortWithError(http.StatusInternalServerError, err)
//			return
//		}
//		if err = global.ExpireRedis(key, time.Hour); err != nil {
//			log.Printf("redis缓存失败：%s", err)
//			c.AbortWithError(http.StatusInternalServerError, err)
//			return
//		}
//
//		// 获取当前IP在 path 上的请求次数
//		accessTime := global.GetLimitRedis(key)
//
//		if err != nil {
//			// 记录日志
//			fmt.Println("get error:", err)
//			c.AbortWithStatus(http.StatusInternalServerError)
//			return
//		}
//		//ip一小时内访问路径超过次数限制，拒绝访问
//		if accessTime > 60 {
//			requestLimit := fmt.Sprintf("request:%s:%s", ip, path)
//			if err = global.RpushRedis(global.InterceptPrefix, requestLimit); err != nil {
//				c.AbortWithStatusJSON(http.StatusServiceUnavailable, err)
//				return
//			}
//			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
//			return
//		}
//		mu.Lock()
//		_, ok := visitorMap[ip]
//		var limiter = rate.NewLimiter(1, 10) // 设置限制为1个请求/秒，最多允许10个并发请求
//		// 如果该IP地址不存在，则创建一个速率限制器
//		if !ok {
//			visitorMap[ip] = limiter
//		}
//		mu.Unlock()
//		// 尝试获取令牌，如果没有可用的令牌则阻塞
//		if !limiter.Allow() {
//			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
//			return
//		}
//		c.Next()
//	}
//}
