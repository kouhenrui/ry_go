package inter

import (
	"errors"
	"gorm.io/gorm"
	"ry_go/src/msg"
	"ry_go/src/pojo"
)

type RoleRepositoryImpl struct {
	roleRepository *pojo.Role
}

type RoleInter interface {
	FindById(id int) (*pojo.Role, error)
}

// TODO 通过ID查找
func (r *RoleRepositoryImpl) FindById(id int) (*pojo.Role, error) {
	r.roleRepository.ID = uint(id)
	err := db.First(r).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(msg.NOT_FOUND_ERROR)
		}
		return nil, err
	}
	return r.roleRepository, nil
}
