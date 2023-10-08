package common

import (
	"mime/multipart"
	"ry_go/src/dto/reqDto"
	util "ry_go/src/utils"
)

var (
	err    error
	capcha = util.NewCaptchaService()
)

/**
 * @Author Khr
 * @Description //TODO 调用方法生成验证码
 * @Date 9:32 2023/9/27
 * @Param
 * @return
 **/
func GetCaptcha() (error, reqDto.Captcha) {

	var newCaptcha reqDto.Captcha

	newCaptcha.Id, newCaptcha.Content, err = capcha.GenerateCaptcha()
	if err != nil {
		return err, newCaptcha
	}
	return nil, newCaptcha
}

/**
 * @Author Khr
 * @Description //TODO 验证输入码
 * @Date 9:33 2023/9/27
 * @Param
 * @return
 **/
func VfCaptcha(capt reqDto.Captcha) bool {
	return capcha.VerifyCaptcha(capt.Id, capt.Content)
}

/**
 * @Author Khr
 * @Description //TODO 文件上穿
 * @Date 9:33 2023/9/27
 * @Param
 * @return
 **/
func UploadFile(file *multipart.FileHeader) {

	//fmt.Println(file, "上传的文件")
	//获取上传文件的类型
	//filetype := file.Header.Get("Content-Type")
	//types := strings.Split(filetype, "/")
	//if types[0] != "image" {
	//	return errors.New(util.FILE_TYPE_ERROR), ""
	//	//c.Error(errors.New(util.FILE_TYPE_ERROR))
	//
	//}
	//name := time.Now().Unix()
	//filename := file.Filename
	//suffix := strings.Split(filename, ".")
	//nameSuffix := suffix[1]
	//t := util.ExistIn(nameSuffix, global.PictureType)
	//if !t {
	//	c.Error(errors.New(util.FILE_SUFFIX_ERROR))
	//	return
	//}

	//newFileName := util.GenerateUniqueFileName(filename)
	//filepath := filepath.Join(global.FilePath, newFileName)
	//fmt.Println("文件路径", filepath)
	//out, err := os.Create(filepath)
	//if err != nil {
	//	return errors.New(util.RESOURCE_NOT_FOUND_ERROR), ""
	//	//c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建文件"})
	//	//return
	//}
	//defer out.Close()
	//
	//io.Copy(out, file)
	//return nil, ""
	// 复制文件内容到目标文件
	//_, err = io.Copy(out, file)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "无法保存文件"})
	//	return
	//}
	//file.Filename = strconv.FormatInt(name, 10) + "." + nameSuffix
	//
	//filePath := path.Join(global.FilePath)
	//_, e := os.Stat(filePath)
	//if e != nil {
	//	os.Mkdir(global.FilePath, os.ModePerm)
	//}
	//pa := path.Join(global.FilePath+"/", file.Filename)
	//io.Copy(filePath, file)
	//c.SaveUploadedFile(file, pa)
	//res.Success(file.Filename)
	//return
}
