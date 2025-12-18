package domain

import (
	"fmt"
	"time"
)

type Student struct {
	ID      int
	Name    string
	Rating  float64
	Mark    int
	AddedAt time.Time
}

// NewStudent создает нового студента
func NewStudent(id int, name string, rating float64) *Student {
	return &Student{
		ID:      id,
		Name:    name,
		Rating:  rating,
		Mark:    0,
		AddedAt: time.Now(),
	}
}

// SetMark устанавливает оценку
func (s *Student) SetMark(mark int) {
	if mark >= 0 && mark <= 10 {
		s.Mark = mark
	}
}

// Compare сравнивает двух студентов для очереди
func (s *Student) Compare(other *Student) int {
	if s.Rating > other.Rating {
		return -1 // выше рейтинг = выше приоритет
	} else if s.Rating < other.Rating {
		return 1
	}

	// При равном рейтинге - FIFO
	if s.AddedAt.Before(other.AddedAt) {
		return -1
	} else if s.AddedAt.After(other.AddedAt) {
		return 1
	}
	return 0
}

// String возвращает строковое представление
func (s *Student) String() string {
	return fmt.Sprintf("%d. %s (рейтинг: %.1f, оценка: %d)",
		s.ID, s.Name, s.Rating, s.Mark)
}
