package repository

import (
	"github.com/mzulfanw/go-unit-testing/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) GetAll() []*entity.Student {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return nil
	} else {
		// Use type assertion to convert the interface{} to []entity.Student
		students := arguments.Get(0).([]entity.Student)

		// Convert []entity.Student to []*entity.Student
		var result []*entity.Student
		for i := range students {
			result = append(result, &students[i])
		}

		return result
	}
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Student {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		student := arguments.Get(0).(entity.Student)
		return &student
	}
}
