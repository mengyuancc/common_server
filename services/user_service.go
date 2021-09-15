package services

import (
	"iris_server/dao"
	"iris_server/datasource"
	"iris_server/models"
)

type UserService interface {
	GetAll() []models.UserInfo
	Search(name string) []models.UserInfo
	Get(id int) *models.UserInfo
	Delete(id int) error
	Update(user *models.UserInfo, columns []string) error
	Create(user *models.UserInfo) error
}

type userService struct {
	dao *dao.UserDao
}

func NewUserService() UserService {
	return &userService{
		dao: dao.NewUserDao(datasource.InstanceMaster()),
	}
}

func (s *userService)GetAll() []models.UserInfo {
	return s.dao.GetAll()
}

func (s *userService)Search(name string) []models.UserInfo {
	return s.dao.Search(name)
}

func (s *userService)Get(id int) *models.UserInfo {
	return s.dao.Get(id)
}
func (s *userService)Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *userService)Update(user *models.UserInfo, columns []string) error {
	return s.dao.Update(user, columns)
}
func (s *userService)Create(user *models.UserInfo) error {
	return s.dao.Create(user)
}
