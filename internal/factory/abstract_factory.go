package factory

import "go/sample-task/internal/repository"

// RepositoryFactory абстрактная фабрика
type RepositoryFactory interface {
	CreateStudentRepository() repository.StudentRepository
}

// InMemoryRepositoryFactory фабрика для репозиториев в памяти
type InMemoryRepositoryFactory struct{}

// NewInMemoryRepositoryFactory создает фабрику
func NewInMemoryRepositoryFactory() RepositoryFactory {
	return &InMemoryRepositoryFactory{}
}

// CreateStudentRepository создает репозиторий студентов в памяти
func (f *InMemoryRepositoryFactory) CreateStudentRepository() repository.StudentRepository {
	return repository.NewInMemoryStudentRepository()
}

// FileRepositoryFactory фабрика для файловых репозиториев
type FileRepositoryFactory struct {
	filePath string
}

// NewFileRepositoryFactory создает фабрику файловых репозиториев
func NewFileRepositoryFactory(filePath string) RepositoryFactory {
	return &FileRepositoryFactory{filePath: filePath}
}

// CreateStudentRepository создает файловый репозиторий
func (f *FileRepositoryFactory) CreateStudentRepository() repository.StudentRepository {
	return repository.NewFileStudentRepository(f.filePath)
}
