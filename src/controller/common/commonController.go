package common

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"ry_go/src/dto/reqDto"
	"ry_go/src/global"
	"ry_go/src/service/common"
	util "ry_go/src/utils"
	"strings"
)

/**
 * @Author Khr
 * @Description //TODO 生成图片验证码
 * @Date 9:21 2023/9/27
 * @Param
 * @return
 **/
func GetCaptcha(c *gin.Context) {
	err, captera := common.GetCaptcha()
	//fmt.Println("接获获取", captera)
	if err != nil {
		c.Error(err)
		return
	}

	c.Set("res", captera)

}

/**
 * @Author Khr
 * @Description //TODO 校验验证码
 * @Date 9:21 2023/9/27
 * @Param
 * @return
 **/
func VfCaptcha(c *gin.Context) {
	var vf = reqDto.Captcha{}
	if err := c.ShouldBindJSON(&vf); err != nil {
		c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &vf))
		return
	}
	c.Set("res", common.VfCaptcha(vf))
	//return common.VfCaptcha(vf)
}

/*
* @MethodName upload
* @Description 上传单个图片,返回字符串
* @Author khr
* @Date 2023/5/8 11:02
 */

func UploadFile(c *gin.Context) {
	//err := c.Request.ParseMultipartForm(10 << 20)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.Error(err)
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.Error(errors.New(util.FILE_TYPE_ERROR))
		//c.Error(errors.New(util.FILE_TYPE_ERROR))
	}
	filename := file.Filename
	newFileName := util.GenerateUniqueFileName(filename)
	filePath := filepath.Join(global.FilePath, newFileName)

	c.SaveUploadedFile(file, filePath)
	c.Set("res", filePath)

}

/*
 * @MethodName uploadVideo
 * @Description 上传视频
 * @Author khr
 * @Date 2023/5/8 11:03
 */

func uploadVideo(c *gin.Context) {
	//if err := c.Request.ParseMultipartForm(10 << 100); err != nil {
	//	c.Error(err)
	//}
	file, _ := c.FormFile("video")
	//获取上传文件的类型
	filetype := file.Header.Get("Content-Type")
	types := strings.Split(filetype, "/")
	fmt.Println(types, "文件类型")
	if types[0] != "video" {
		c.Error(errors.New(util.FILE_TYPE_ERROR))
	}
	filename := file.Filename
	newFileName := util.GenerateUniqueFileName(filename)
	filePath := filepath.Join(global.FilePath, newFileName)
	//name := time.Now().Unix()
	//filename := file.Filename
	//suffix := strings.Split(filename, ".")
	//nameSuffix := suffix[1]
	//t := util.ExistIn(nameSuffix, global.VideoType)
	//if !t {
	//	res.Error(http.StatusBadRequest, util.FILE_SUFFIX_ERROR)
	//	return
	//}
	//file.Filename = strconv.FormatInt(name, 10) + "." + nameSuffix
	//filePath := path.Join(global.VideoPath)
	//_, e := os.Stat(filePath)
	//if e != nil {
	//	os.Mkdir(global.VideoPath, os.ModePerm)
	//}
	//pa := path.Join("./"+global.VideoPath+"/", file.Filename)
	c.SaveUploadedFile(file, filePath)
	c.Set("res", filePath)
}
