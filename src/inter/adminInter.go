package inter

import (
	"errors"
	"gorm.io/gorm"
	"ry_go/src/dto/reqDto"
	"ry_go/src/dto/resDto"
	"ry_go/src/global"
	"ry_go/src/pojo"
	"time"
)

var (
	db           = global.MysqlDClient
	a            = &pojo.Admin{}
	resAdminList []resDto.AdminList
	count        int64
	reslist      resDto.CommonList
)

type AdminRepositoryImpl struct{}

type AdminRepositoryInter interface {
	AdminList(list *reqDto.AdminList) (*resDto.CommonList, error)
	CheckByName(userName string) (*pojo.Admin, error)
	CheckByPhone(phone string) (*pojo.Admin, error)
	CheckByNickName(nickName string) (*pojo.Admin, error)
	RemoveAccessToken(id uint) error
	UpdateToken(access_token string, id uint, ip string) error
	AddAdmin(body pojo.Admin) error
}

// TODO 根据手机号查询
func (ap *AdminRepositoryImpl) CheckByPhone(phone string) (*pojo.Admin, error) {
	a.Phone = phone
	err := db.Preload("Role").First(a).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return a, nil
}

// TODO 列表
func (ap *AdminRepositoryImpl) AdminList(list *reqDto.AdminList) (*resDto.CommonList, error) {
	query := db.Model(a)
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

// TODO 查询账号
func (ap *AdminRepositoryImpl) CheckByNickName(nickName string) (*pojo.Admin, error) {
	a.NickName = nickName
	err := db.Preload("Role").First(a).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return a, nil
}

// TODO 查询昵称
func (ap *AdminRepositoryImpl) CheckByName(userName string) (*pojo.Admin, error) {
	a.UserName = userName
	err := db.First(a).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err

	}
	return a, nil
}

// TODO 详情数据
func (ap *AdminRepositoryImpl) AdminInfo(id uint) (*resDto.AdminInfo, error) {
	a.ID = id
	err := db.Model(a).Preload("Role").Find(resDto.AdminInfo{}).Error //.Select("inter.*,r.name as role_name").Joins("left join rule as r on r.id = inter.role").Scan(&adminInfo).Error
	if err != nil {
		return nil, err
	}
	return &resDto.AdminInfo{}, nil
}

// TODO 更新token数据
func (ap *AdminRepositoryImpl) UpdateToken(access_token string, id uint, ip string) error {
	a.ID = id
	err := db.Model(a).Updates(pojo.Admin{AccessToken: access_token, LoginAt: time.Now(), LoginIp: ip}).Error
	if err != nil {
		return err
	}
	return nil
}

// TODO 去除token
func (ap *AdminRepositoryImpl) RemoveAccessToken(id uint) error {
	a.ID = id
	return db.Model(a).Update("access_token", "").Error

}

// TODO 增加用户
func (ap *AdminRepositoryImpl) AddAdmin(admin pojo.Admin) error {
	return db.Create(&admin).Error
}
