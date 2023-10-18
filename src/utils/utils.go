package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"path/filepath"
	"reflect"
	"regexp"
	"ry_go/src/msg"
	"strconv"
	"time"
)

// json格式化数据
func Marshal(user interface{}) []byte {
	ub, _ := json.Marshal(user)
	return ub
}
func UnMarshal(r []byte, res interface{}) (bool, interface{}) {
	err := json.Unmarshal(r, &res)
	if err != nil {
		return false, msg.REDIS_INFORMATION_ERROR
	}
	return true, res
}

/*
 * @MethodName 参数验证
 * @Description
 * @Author khr
 * @Date 2023/8/21 10:21
 */
func GetValidate(err error, obj any) error {

	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		fmt.Println("param error:", invalid)
		return invalid
	}
	//反射获取标签的注释
	getObj := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {
		return errs
		for _, e := range errs {
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				msg := f.Tag.Get("msg")
				return errors.New(msg)
			}
		}
	}
	return err
}

/*
 * @MethodName ExistIn
 * @Description 判断参数是否存在
 * @Author khr
 * @Date 2023/4/14 8:52
 */

func ExistIn(param string, paths []string) bool {
	for _, v := range paths {
		if param == v {
			return true
		}
	}
	return false
}

/*
 * @MethodName FuzzyMatch
 * @Description 正则模糊匹配路径
 * @Author khr
 * @Date 2023/5/9 16:25
 */
func FuzzyMatch(param string, paths []string) bool {
	for _, y := range paths {
		if regexp.MustCompile(y).MatchString(param) {

			//fmt.Print("匹配道路进了")
			return true
		}

	}
	return false
}

/**
 * @Author Khr
 * @Description //TODO 根据时间戳生成名称
 * @Date 15:41 2023/9/27
 * @Param
 * @return
 **/
func GenerateUniqueFileName(originalFileName string) string {
	// 生成唯一的文件名，可以使用时间戳或随机数等方式
	timestamp := time.Now().UnixNano()
	extension := filepath.Ext(originalFileName)
	uniqueFileName := strconv.FormatInt(timestamp, 10) + extension
	return uniqueFileName
}

/**
 * @Author Khr
 * @Description //TODO 文件流下载
 * @Date 16:48 2023/10/16
 * @Param
 * @return
 **/
//func DownLoadFileHandler(w http.ResponseWriter, req *http.Request, fpath string, resPrefix string) {
//	defer req.Body.Close()
//	err := req.ParseForm()
//	//request header
//	ranStr := req.Header.Get("Range")
//	fmt.Println("Range: " + ranStr)
//	//    w.Header().Set("Access-Control-Allow-Origin", "*")
//	if err != nil {
//		w.Header().Set("Content-Type", " text/plain; charset=UTF-8")
//		w.WriteHeader(http.StatusBadRequest)
//		w.Write([]byte(err.Error()))
//		return
//	}
//	if fpath == "" || !strings.HasPrefix(fpath, resPrefix) {
//		w.Header().Set("Content-Type", " text/plain; charset=UTF-8")
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte("No file can be downloaded."))
//		return
//	}
//	file, err := os.Open(fpath)
//	if err != nil {
//		w.Header().Set("Content-Type", " text/plain; charset=UTF-8")
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(err.Error()))
//		return
//	}
//	defer file.Close()
//	//设置下载完成以后的名称
//	aliasName := path.Base(fpath)
//	w.Header().Set("Content-Disposition", "attachment; filename="+aliasName)
//	//设置下载的偏移量
//	fstat, _ := file.Stat()
//	fsize := fstat.Size()
//	var spos int64
//	if ranStr != "" {
//		rs := strings.Split(strings.TrimPrefix(ranStr, "bytes="), ",")[0]
//		sePos := strings.Split(rs, "-")
//		spStr := sePos[0]
//		spos, _ = strconv.ParseInt(spStr, 0, 64)
//		file.Seek(spos, 0)
//	}
//	fmt.Println(spos)
//	w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", spos, fsize-1, fsize))
//	w.Header().Set("Content-Length", fmt.Sprintf("%d", fsize-spos))
//	//    w.WriteHeader(http.StatusPartialContent)
//	if spos == 0 {
//		w.WriteHeader(http.StatusOK)
//	} else {
//		w.WriteHeader(http.StatusPartialContent)
//	}
//	wcount, err := io.Copy(w, file)
//	if err != nil {
//		fmt.Println("io.Copy: ", err.Error())
//	}
//	if spos+wcount == fsize {
//		fmt.Println("remove from cache info.")
//		fmt.Println("delete finished file." + fpath)
//		file.Close()
//		err := os.Remove(fpath)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//}
