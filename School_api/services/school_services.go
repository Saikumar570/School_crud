package services

import (
	"School_api/models"
	"School_api/repositories"
)

type SchoolService interface {
	GetSchools() ([]models.School, error)
	GetSchoolByID(id uint) (models.School, error)
	CreateSchool(school models.School) (models.School, error)
	UpdateSchool(id uint, school models.School) (models.School, error)
	DeleteSchool(id uint) error
}

type schoolService struct {
	repo repositories.SchoolRepository
}

func NewSchoolService(repo repositories.SchoolRepository) SchoolService {
	return &schoolService{repo}
}

func (s *schoolService) GetSchools() ([]models.School, error) {
	return s.repo.FindAll()
}

func (s *schoolService) GetSchoolByID(id uint) (models.School, error) {
	return s.repo.FindByID(id)
}

func (s *schoolService) CreateSchool(school models.School) (models.School, error) {
	return s.repo.Create(school)
}

func (s *schoolService) UpdateSchool(id uint, school models.School) (models.School, error) {
	existingSchool, err := s.repo.FindByID(id)
	if err != nil {
		return existingSchool, err
	}
	existingSchool.Name = school.Name
	existingSchool.SchoolId = school.SchoolId
	existingSchool.SchoolAddress = school.SchoolAddress
	existingSchool.Classes = school.Classes
	return s.repo.Update(existingSchool)
}

func (s *schoolService) DeleteSchool(id uint) error {
	school, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(school)
}
