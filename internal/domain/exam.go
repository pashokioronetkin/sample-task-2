package domain

import (
	"container/heap"
	"fmt"
)

// Item представляет элемент в очереди
type Item struct {
	student   *Student
	priority  float64 // рейтинг студента
	index     int     // индекс в куче
	timestamp int64   // время добавления для FIFO
}

// PriorityQueue реализует очередь с приоритетами
type PriorityQueue struct {
	items   []*Item
	counter int64
}

// NewPriorityQueue создает новую очередь
func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{
		items:   make([]*Item, 0),
		counter: 0,
	}
	heap.Init(pq)
	return pq
}

// Len возвращает количество элементов
func (pq *PriorityQueue) Len() int {
	return len(pq.items)
}

// Less определяет порядок сортировки
func (pq *PriorityQueue) Less(i, j int) bool {
	if pq.items[i].priority != pq.items[j].priority {
		return pq.items[i].priority > pq.items[j].priority // выше рейтинг = выше приоритет
	}
	// При равном рейтинге - кто раньше добавлен
	return pq.items[i].timestamp < pq.items[j].timestamp
}

// Swap меняет элементы местами
func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].index = i
	pq.items[j].index = j
}

// Push добавляет элемент
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(pq.items)
	item := x.(*Item)
	item.index = n
	pq.items = append(pq.items, item)
}

// Pop удаляет и возвращает элемент с наивысшим приоритетом
func (pq *PriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	item.index = -1
	pq.items = old[0 : n-1]
	return item
}

// Enqueue добавляет студента в очередь
func (pq *PriorityQueue) Enqueue(student *Student) {
	pq.counter++
	item := &Item{
		student:   student,
		priority:  student.Rating,
		timestamp: pq.counter,
	}
	heap.Push(pq, item)
}

// Dequeue извлекает студента с наивысшим приоритетом
func (pq *PriorityQueue) Dequeue() *Student {
	if pq.Len() == 0 {
		return nil
	}
	item := heap.Pop(pq).(*Item)
	return item.student
}

// IsEmpty проверяет, пуста ли очередь
func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Len() == 0
}

// Size возвращает размер очереди
func (pq *PriorityQueue) Size() int {
	return pq.Len()
}

// Exam представляет экзамен
type Exam struct {
	Name    string
	queue   *PriorityQueue
	results []*Student
}

// NewExam создает новый экзамен
func NewExam(name string) *Exam {
	return &Exam{
		Name:    name,
		queue:   NewPriorityQueue(),
		results: make([]*Student, 0),
	}
}

// AddStudent добавляет студента на экзамен
func (e *Exam) AddStudent(student *Student) {
	e.queue.Enqueue(student)
}

// Conduct проводит экзамен
func (e *Exam) Conduct() {
	fmt.Printf("Начало экзамена: %s\n", e.Name)
	fmt.Println("Порядок сдачи по рейтингу:")

	position := 1
	for !e.queue.IsEmpty() {
		student := e.queue.Dequeue()

		// Рассчитываем оценку на основе рейтинга
		mark := e.calculateMark(student.Rating)
		student.SetMark(mark)
		e.results = append(e.results, student)

		fmt.Printf("%d. %s сдает... Получает оценку: %d\n",
			position, student.Name, mark)
		position++
	}

	fmt.Printf("Экзамен завершен. Всего сдало: %d студентов\n", len(e.results))
}

// calculateMark рассчитывает оценку на основе рейтинга
func (e *Exam) calculateMark(rating float64) int {
	switch {
	case rating >= 90:
		return 10
	case rating >= 80:
		return 9
	case rating >= 70:
		return 8
	case rating >= 60:
		return 7
	case rating >= 50:
		return 6
	default:
		return 5
	}
}

// GetQueueSize возвращает количество студентов в очереди
func (e *Exam) GetQueueSize() int {
	return e.queue.Size()
}

// GetResults возвращает результаты экзамена
func (e *Exam) GetResults() []*Student {
	return e.results
}

// String возвращает информацию об экзамене
func (e *Exam) String() string {
	return fmt.Sprintf("Экзамен '%s', студентов: %d",
		e.Name, e.GetQueueSize())
}
