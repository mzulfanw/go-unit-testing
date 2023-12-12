package service

import (
	"errors"
	"github.com/mzulfanw/go-unit-testing/entity"
	"github.com/mzulfanw/go-unit-testing/repository"
)

type StudentService struct {
	Repository repository.StudentRepository
}

func (service StudentService) GetAll() ([]*entity.Student, error) {
	students := service.Repository.GetAll()

	if students == nil {
		return nil, errors.New("empty Student")
	} else {
		return students, nil
	}
}

func (service StudentService) FindById(id string) (*entity.Student, error) {
	student := service.Repository.FindById(id)

	if student == nil {
		return nil, errors.New("not found Student")
	} else {
		return student, nil
	}
}
