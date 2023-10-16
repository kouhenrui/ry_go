package pojo

import (
	"gorm.io/gorm"
	"log"
)

func Repositoryinit(db *gorm.DB) {
	err := db.AutoMigrate(
		&Admin{}, //管理员表
		&Role{},  //角色表
	)
	if err != nil {
		panic(err)
	}
	log.Println("结构表创建成功")
}
