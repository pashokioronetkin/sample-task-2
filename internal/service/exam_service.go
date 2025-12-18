package service

import (
	"fmt"
	"go/sample-task/internal/domain"
	"go/sample-task/internal/repository"
)

// ExamService сервис для управления экзаменами
type ExamService struct {
	studentRepo repository.StudentRepository
}

// NewExamService создает новый сервис экзаменов
func NewExamService(repo repository.StudentRepository) *ExamService {
	return &ExamService{studentRepo: repo}
}

// ConductExam проводит экзаменационную сессию
func (s *ExamService) ConductExam(examName string) (*domain.Exam, error) {
	// Загружаем студентов из репозитория
	students, err := s.studentRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("ошибка загрузки студентов: %v", err)
	}

	if len(students) == 0 {
		return nil, fmt.Errorf("нет студентов для экзамена")
	}

	fmt.Printf("Загружено %d студентов\n", len(students))

	// Создаем экзамен
	exam := domain.NewExam(examName)

	// Добавляем студентов в очередь экзамена
	for _, student := range students {
		exam.AddStudent(student)
	}

	fmt.Printf("Сформирована очередь из %d студентов\n", exam.GetQueueSize())
	fmt.Println("Порядок сдачи по рейтингу:")

	// Проводим экзамен
	exam.Conduct()

	// Сохраняем результаты
	results := exam.GetResults()
	for _, student := range results {
		if err := s.studentRepo.Save(student); err != nil {
			fmt.Printf("Ошибка сохранения студента %s: %v\n", student.Name, err)
		}
	}

	fmt.Println("Результаты сохранены")
	return exam, nil
}

// PrintStatistics выводит статистику по экзамену
func (s *ExamService) PrintStatistics(exam *domain.Exam) {
	results := exam.GetResults()

	if len(results) == 0 {
		fmt.Println("Нет результатов для статистики")
		return
	}

	var totalRating, totalMark float64
	var passed, failed int

	for _, student := range results {
		totalRating += student.Rating
		totalMark += float64(student.Mark)

		if student.Mark >= 6 {
			passed++
		} else {
			failed++
		}
	}

	fmt.Println("\n=== Статистика экзамена ===")
	fmt.Printf("Всего студентов: %d\n", len(results))
	fmt.Printf("Сдали успешно: %d\n", passed)
	fmt.Printf("Не сдали: %d\n", failed)
	fmt.Printf("Средний рейтинг: %.1f\n", totalRating/float64(len(results)))
	fmt.Printf("Средняя оценка: %.1f\n", totalMark/float64(len(results)))
}
