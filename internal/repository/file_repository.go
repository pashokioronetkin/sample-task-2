package repository

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go/sample-task/internal/domain"
)

// FileStudentRepository реализация для работы с файлом
type FileStudentRepository struct {
	filePath string
	students map[int]*domain.Student
}

// NewFileStudentRepository создает файловый репозиторий
func NewFileStudentRepository(filePath string) *FileStudentRepository {
	return &FileStudentRepository{
		filePath: filePath,
		students: make(map[int]*domain.Student),
	}
}

// GetAll загружает студентов из файла
func (r *FileStudentRepository) GetAll() ([]*domain.Student, error) {
	// Читаем файл
	file, err := os.Open(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Файла нет - возвращаем пустой список
			return []*domain.Student{}, nil
		}
		return nil, fmt.Errorf("ошибка открытия файла: %v", err)
	}
	defer file.Close()

	r.students = make(map[int]*domain.Student)
	var students []*domain.Student

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue // Пропускаем пустые строки и комментарии
		}

		student, err := r.parseStudentLine(line)
		if err != nil {
			fmt.Printf("Ошибка в строке %d: %v\n", lineNum, err)
			continue
		}

		r.students[student.ID] = student
		students = append(students, student)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %v", err)
	}

	return students, nil
}

// Save сохраняет студента
func (r *FileStudentRepository) Save(student *domain.Student) error {
	r.students[student.ID] = student
	return r.saveToFile()
}

// SaveAll сохраняет всех студентов
func (r *FileStudentRepository) SaveAll(students []*domain.Student) error {
	for _, student := range students {
		r.students[student.ID] = student
	}
	return r.saveToFile()
}

// FindByID ищет студента по ID
func (r *FileStudentRepository) FindByID(id int) (*domain.Student, error) {
	student, exists := r.students[id]
	if !exists {
		return nil, fmt.Errorf("студент с ID %d не найден", id)
	}
	return student, nil
}

// parseStudentLine парсит строку файла
func (r *FileStudentRepository) parseStudentLine(line string) (*domain.Student, error) {
	parts := strings.Split(line, ",")
	if len(parts) < 3 {
		return nil, fmt.Errorf("неверный формат, ожидается: id,name,rating[,mark]")
	}

	// ID
	id, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return nil, fmt.Errorf("неверный ID: %v", err)
	}

	// Имя
	name := strings.TrimSpace(parts[1])
	if name == "" {
		return nil, fmt.Errorf("имя не может быть пустым")
	}

	// Рейтинг
	rating, err := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)
	if err != nil {
		return nil, fmt.Errorf("неверный рейтинг: %v", err)
	}

	// Оценка (опционально)
	mark := 0
	if len(parts) > 3 {
		mark, err = strconv.Atoi(strings.TrimSpace(parts[3]))
		if err != nil {
			mark = 0
		}
	}

	student := domain.NewStudent(id, name, rating)
	student.SetMark(mark)

	return student, nil
}

// saveToFile сохраняет всех студентов в файл
func (r *FileStudentRepository) saveToFile() error {
	file, err := os.Create(r.filePath)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Записываем заголовок
	_, err = writer.WriteString("# Формат: ID,Имя,Рейтинг,Оценка\n")
	if err != nil {
		return err
	}

	// Записываем студентов
	for _, student := range r.students {
		line := fmt.Sprintf("%d,%s,%.1f,%d\n",
			student.ID, student.Name, student.Rating, student.Mark)
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
