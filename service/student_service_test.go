package service

import (
	"github.com/mzulfanw/go-unit-testing/entity"
	"github.com/mzulfanw/go-unit-testing/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var studentRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var studentService = StudentService{Repository: studentRepository}

func TestStudentService_GetAll(t *testing.T) {
	students := []entity.Student{
		{
			Id:      "1",
			Name:    "Muhammad Zulfan Wahyudin",
			Age:     21,
			Address: "Kolmas",
		},
		{
			Id:      "2",
			Name:    "Jono",
			Age:     20,
			Address: "Pakusarakan Kuda",
		},
	}

	studentRepository.Mock.On("GetAll").Return(students)

	result, err := studentService.GetAll()

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestStudentService_FindById(t *testing.T) {
	student := entity.Student{
		Id:      "1",
		Name:    "Muhammad Zulfan Wahyudin",
		Age:     21,
		Address: "Kolmas",
	}

	studentRepository.Mock.On("FindById", "1").Return(student)

	result, err := studentService.FindById("1")

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
