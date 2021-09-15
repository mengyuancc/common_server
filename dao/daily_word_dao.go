package dao

import (
	"github.com/go-xorm/xorm"
	"iris_server/models"
)

type DailyWordDao struct {
	engine *xorm.Engine
}

func NewDailyWordDao(engine *xorm.Engine) *DailyWordDao {
	return &DailyWordDao{
		engine:engine,
	}
}

func (d *DailyWordDao) Get(id int) *models.DailyWord {
	data := &models.DailyWord{Id:id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *DailyWordDao) GetAll() []models.DailyWord {
	datalist := make([]models.DailyWord, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *DailyWordDao) Search(name string) []models.DailyWord {
	datalist := make([]models.DailyWord, 0)
	err := d.engine.Where("name=?", name).
		Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *DailyWordDao) Delete(id int) error {
	data := &models.DailyWord{Id:id}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *DailyWordDao) Update(data *models.DailyWord, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *DailyWordDao) Create(data *models.DailyWord) error {
	_, err := d.engine.Insert(data)
	return err
}
