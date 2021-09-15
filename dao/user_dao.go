package dao

import (
	"github.com/go-xorm/xorm"
	. "iris_server/models"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{
		engine:engine,
	}
}

func (d *UserDao) Get(id int) *UserInfo {
	data := &UserInfo{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *UserDao) GetAll() []UserInfo {
	datalist := make([]UserInfo, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *UserDao) Search(name string) []UserInfo {
	datalist := make([]UserInfo, 0)
	err := d.engine.Where("name=?", name).
		Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *UserDao) Delete(id int) error {
	data := &UserInfo{Id: id}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *UserDao) Update(data *UserInfo, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *UserDao) Create(data *UserInfo) error {
	_, err := d.engine.Insert(data)
	return err
}
