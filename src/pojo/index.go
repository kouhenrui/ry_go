package pojo

import (
	"fmt"
	"log"
	"ry_go/src/dto/resDto"
	"ry_go/src/global"
)

var (
	db = global.MysqlDClient

	reslist = resDto.CommonList{}
	count   int64
)

func Repositoryinit() {
	db.AutoMigrate(&Admin{})
	log.Fatalln("结构表创建成功")
	fmt.Println("结构表创建成功")
}
