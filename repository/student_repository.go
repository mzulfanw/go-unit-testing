package repository

import "github.com/mzulfanw/go-unit-testing/entity"

type StudentRepository interface {
	GetAll() []*entity.Student
	FindById(id string) *entity.Student
}
