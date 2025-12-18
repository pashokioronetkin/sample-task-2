package main

// Задание 2
// 1.	Статический (Ad hoc) полиморфизм
// 2.	Спроектировать класс Студент с атрибутами: id, name, rating, mark и класс "Экзамен" на основе очереди с приоритетами. Построить контекстную диаграмму для Преподавателя, Студента, Заведующего отделением
// 3.	Написать программу для демонстрации экзамена. Для хранения данных использовать текстовый файл. Студент с более высоким рейтингом должен встать в начало очереди. При равенстве рейтингов первым используется FIFO
// 4.	Порождающие паттерны проектирования. Абстрактная фабрика (Abstract Factory)

// Решение предоставил студент группы 402ИС-22 Донских П.

import (
	"fmt"
	"go/sample-task/internal/domain"
	"go/sample-task/internal/factory"
	"go/sample-task/internal/polymorphism"
	"go/sample-task/internal/service"
	"strings"
)

func main() {
	fmt.Print("=== СИСТЕМА УПРАВЛЕНИЯ ЭКЗАМЕНАМИ ===\n")

	// 1. Демонстрация Ad hoc полиморфизма
	polymorphism.Demo()

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	// 2. Создаем фабрику для файлового репозитория (Abstract Factory)
	repoFactory := factory.NewFileRepositoryFactory("data/students.txt")
	studentRepo := repoFactory.CreateStudentRepository()

	// 3. Загружаем студентов из файла
	fmt.Println("Загрузка студентов из файла data/students.txt...")
	students, err := studentRepo.GetAll()
	if err != nil {
		fmt.Printf("Ошибка загрузки студентов: %v\n", err)
		return
	}

	if len(students) == 0 {
		fmt.Println("В файле нет студентов. Добавляем тестовых...")
		students = []*domain.Student{
			domain.NewStudent(1, "Иванов Иван", 85.5),
			domain.NewStudent(2, "Петров Петр", 92.0),
			domain.NewStudent(3, "Сидорова Анна", 78.3),
			domain.NewStudent(4, "Козлов Дмитрий", 92.0),
			domain.NewStudent(5, "Михайлова Елена", 88.7),
		}
		for _, student := range students {
			studentRepo.Save(student)
		}
	} else {
		fmt.Printf("Загружено %d студентов из файла\n", len(students))
	}

	// 4. Создаем сервис и проводим экзамен
	examService := service.NewExamService(studentRepo)

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("=== ПРОВЕДЕНИЕ ЭКЗАМЕНА ===")

	exam, err := examService.ConductExam("Программирование на Go")
	if err != nil {
		fmt.Printf("Ошибка проведения экзамена: %v\n", err)
		return
	}

	// 5. Выводим статистику
	examService.PrintStatistics(exam)

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Экзамен завершен успешно!")
	fmt.Println("Результаты сохранены в data/students.txt")
}
