package repository

import (
	"fmt"
	"go/sample-task/internal/domain"
)

// InMemoryStudentRepository репозиторий в памяти
type InMemoryStudentRepository struct {
	students map[int]*domain.Student
}

// NewInMemoryStudentRepository создает репозиторий в памяти
func NewInMemoryStudentRepository() *InMemoryStudentRepository {
	return &InMemoryStudentRepository{
		students: make(map[int]*domain.Student),
	}
}

// GetAll возвращает всех студентов
func (r *InMemoryStudentRepository) GetAll() ([]*domain.Student, error) {
	var students []*domain.Student
	for _, student := range r.students {
		students = append(students, student)
	}
	return students, nil
}

// Save сохраняет студента
func (r *InMemoryStudentRepository) Save(student *domain.Student) error {
	r.students[student.ID] = student
	return nil
}

// SaveAll сохраняет всех студентов
func (r *InMemoryStudentRepository) SaveAll(students []*domain.Student) error {
	for _, student := range students {
		r.students[student.ID] = student
	}
	return nil
}

// FindByID ищет студента по ID
func (r *InMemoryStudentRepository) FindByID(id int) (*domain.Student, error) {
	student, exists := r.students[id]
	if !exists {
		return nil, fmt.Errorf("студент с ID %d не найден", id)
	}
	return student, nil
}
