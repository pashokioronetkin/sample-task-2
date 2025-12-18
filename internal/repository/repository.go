package repository

import "go/sample-task/internal/domain"

// StudentRepository интерфейс для работы со студентами
type StudentRepository interface {
	GetAll() ([]*domain.Student, error)
	Save(student *domain.Student) error
	SaveAll(students []*domain.Student) error
	FindByID(id int) (*domain.Student, error)
}
