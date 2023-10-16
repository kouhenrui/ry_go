package pojo

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

func Repositoryinit(db *gorm.DB) {
	db.AutoMigrate(
		&Admin{}, //管理员表
		&Role{},  //角色表
	)
	log.Fatalln("结构表创建成功")
	fmt.Println("结构表创建成功")
}
