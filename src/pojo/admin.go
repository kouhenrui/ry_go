package pojo

import (
	"gorm.io/gorm"
	"ry_go/src/dto/reqDto"
	"ry_go/src/dto/resDto"
	"time"
)

type Admin struct {
	gorm.Model
	UserName    string    `json:"user_name" gorm:"not null ;comment:'账号'" `
	NickName    string    `json:"nick_name" gorm:"comment:'昵称'"`
	Password    string    `json:"password" gorm:"comment:'密码'"`
	Salt        string    `json:"salt" gorm:"comment:'加密盐'"`
	Phone       string    `json:"phone" gorm:"unique:true;comment:'手机号'"`
	AccessToken string    `json:"access_token" gorm:"comment:'登录状态'"`
	Sex         int       `json:"sex" gorm:"comment:'性别'"`
	Avatar      string    `json:"avatar" gorm:"comment:'头像'"`
	Email       string    `json:"email" gorm:"comment:'邮箱'"`
	Role        int       `json:"role" gorm:"comment:'权限'"`
	LoginIp     string    `json:"login_ip" gorm:"comment:'登录ip'"`
	LoginAt     time.Time `json:"login_at" gorm:"comment:'登陆时间'"`
}

func AdminServiceImpl() Admin {
	return Admin{}
}

var (
	//admins=[]Admin{}
	adminInfo    = resDto.AdminInfo{}
	resAdminList = []resDto.AdminList{} //要查询的字段
)

// 分页,模糊查询用户
func (a *Admin) AdminList(list reqDto.AdminList) (*resDto.CommonList, error) {
	query := db.Model(&a)
	if list.Name != "" {
		query.Where("user_name like ?", "%"+list.Name+"%")
	}
	err := query.Limit(list.Take).Offset(int(list.Skip)).Find(&resAdminList).Count(&count).Error
	reslist.Count = uint(count)
	reslist.List = resAdminList
	if err != nil {
		return nil, err
	}
	return &reslist, nil
}

// 查询账号
func (a *Admin) CheckByNickName(nickName string) (*Admin, error) {
	a.NickName = nickName
	err := db.First(&a).Error
	if err != nil {
		return nil, err
	}
	return a, nil
}

// 查询名称
func (a *Admin) CheckByName(userName string) (*Admin, error) {
	a.UserName = userName
	err := db.First(&a).Error
	if err != nil {
		return nil, err
	}
	return a, nil
}

// 详情数据
func (a *Admin) AdminInfo(id int) (*resDto.AdminInfo, error) {
	a.ID = uint(id)
	err := db.Model(&a).Select("admin.name, admin.account,admin.role,r.name as role_name").Joins("left join rule as r on r.id = admin.role").Scan(&adminInfo).Error
	if err != nil {
		return nil, err
	}
	return &adminInfo, nil
}

// 更新token数据
func (a *Admin) UpdateToken(access_token string, id uint) error {
	a.ID = id
	err := db.Model(&a).Update("access_token", access_token).Error
	if err != nil {
		return err
	}
	return nil
}

// 增加用户
func (a *Admin) AddAdmin(admins Admin) error {
	err := db.Create(&admins).Error
	if err != nil {
		return err
	}
	return err
}
