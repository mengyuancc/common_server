package services

import (
	"dao"
	"models"
	"datasource"
)

type DailyWordService interface {
	GetAll() []models.DailyWord
	Search(name string) []models.DailyWord
	Get(id int) *models.DailyWord
}

type dailyWordService struct {
	dao *dao.DailyWordDao
}

func NewDailyWordService() DailyWordService {
	return &dailyWordService{
		dao: dao.NewDailyWordDao(datasource.InstanceMaster()),
	}
}

func (s *dailyWordService)GetAll() []models.DailyWord {
	return s.dao.GetAll()
}

func (s *dailyWordService)Search(name string) []models.DailyWord {
	return s.dao.Search(name)
}

func (s *dailyWordService)Get(id int) *models.DailyWord {
	return s.dao.Get(id)
}
