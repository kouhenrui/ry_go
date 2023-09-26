package pojo

import (
	"log"
	"ry_go/src/dto/resDto"
	"ry_go/src/global"
)

var (
	db = global.MysqlDClient

	reslist = resDto.CommonList{}
	count   int64
)

func init() {
	db.AutoMigrate()
	log.Fatalln("结构表创建成功")
}
