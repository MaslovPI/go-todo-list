package datalayer

import "time"

const (
	ErrNotFound         = VaultErr("Could not find task")
	ErrTaskDoesNotExist = VaultErr("cannot perform operation on task because it does not exist")
)

type VaultErr string

func (e VaultErr) Error() string {
	return string(e)
}

type Task struct {
	ID          uint
	Description string
	CreatedAt   time.Time
	IsComplete  bool
}

type TimeProvder interface {
	GetTimeStamp() time.Time
}

type DefaultTimeProvider struct{}

func (d *DefaultTimeProvider) GetTimeStamp() time.Time {
	return time.Now()
}

type TaskVault interface {
	add(task Task)
	get(id uint) (Task, error)
	getNextId() uint
}

type MapTaskVault struct {
	db     map[uint]Task
	lastId uint
}

func (m *MapTaskVault) init() {
	m.db = make(map[uint]Task)
	m.lastId = 0
}

func (m *MapTaskVault) add(task Task) {
	m.db[task.ID] = task
}

func (m *MapTaskVault) get(id uint) (Task, error) {
	task, ok := m.db[id]
	if !ok {
		return Task{}, ErrNotFound
	}
	return task, nil
}

func (m *MapTaskVault) getNextId() uint {
	m.lastId++
	return uint(m.lastId)
}

func AddTask(description string, vault TaskVault, timeProvider TimeProvder) uint {
	id := vault.getNextId()
	task := Task{
		ID:          id,
		Description: description,
		CreatedAt:   timeProvider.GetTimeStamp(),
		IsComplete:  false,
	}
	vault.add(task)
	return id
}

func GetTask(id uint, vault TaskVault) (Task, error) {
	task, err := vault.get(id)
	return task, err
}
