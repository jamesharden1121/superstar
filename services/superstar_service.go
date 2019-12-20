package services

import (
	"superStar/dao"
	"superStar/datasource"
	"superStar/models"
)

type SuperstarService interface {
	Search(country string) []models.StarInfo
	GetAll() []models.StarInfo
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(user *models.StarInfo) error
	Create(starInfo *models.StarInfo) error
}

func NewSuperstarService() SuperstarService {
	return &superstarService{dao: dao.NewSuperstarDao(datasource.InstanceMaster())}
}

type superstarService struct {
	dao *dao.SuperstarDao
}

func (s *superstarService) Search(country string) []models.StarInfo {
	return s.dao.Search(country)
}

func (s *superstarService) GetAll() []models.StarInfo {
	return s.dao.GetAll()
}

func (s *superstarService) Get(id int) *models.StarInfo {
	return s.dao.Get(id)
}

func (s *superstarService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *superstarService) Update(user *models.StarInfo) error {
	return s.dao.Upedate(user)
}

func (s *superstarService) Create(starInfo *models.StarInfo) error {
	return s.dao.Create(starInfo)
}
