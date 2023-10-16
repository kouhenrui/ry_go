package inter

import (
	"errors"
	"gorm.io/gorm"
	"ry_go/src/msg"
	"ry_go/src/pojo"
)

var (
	r = &pojo.Role{}
)

type RoleRepositoryImpl struct {
}

type RoleInter interface {
	FindById(id int) (*pojo.Role, error)
}

// TODO 通过ID查找
func (rp *RoleRepositoryImpl) FindById(id int) (*pojo.Role, error) {
	r.ID = uint(id)
	err := db.First(r).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(msg.NOT_FOUND_ERROR)
		}
		return nil, err
	}
	return r, nil
}
