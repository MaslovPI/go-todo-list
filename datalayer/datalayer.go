package datalayer

import (
	"maps"
	"slices"
	"time"
)

const (
	ErrNotFound             = VaultErr("Could not find task")
	ErrTaskDoesNotExist     = VaultErr("Cannot perform operation on task because it does not exist")
	ErrTaskDescriptionEmpty = VaultErr("Task description cannot be empty")
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

type TimeProvider interface {
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
	list() []Task
	listUnfinished() []Task
}

type MapTaskVault struct {
	db     map[uint]Task
	lastId uint
}

func NewMapTaskVault() *MapTaskVault {
	vault := &MapTaskVault{
		db:     make(map[uint]Task),
		lastId: 0,
	}
	return vault
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

func (m *MapTaskVault) list() []Task {
	return slices.Collect(maps.Values(m.db))
}

func (m *MapTaskVault) listUnfinished() []Task {
	return slices.Collect(func(yield func(Task) bool) {
		for _, task := range slices.Collect(maps.Values(m.db)) {
			if !task.IsComplete {
				if !yield(task) {
					return
				}
			}
		}
	})
}

func (m *MapTaskVault) getNextId() uint {
	m.lastId++
	return m.lastId
}

func AddTask(description string, vault TaskVault, timeProvider TimeProvider) (uint, error) {
	if description == "" {
		return 0, ErrTaskDescriptionEmpty
	}

	id := vault.getNextId()
	task := Task{
		ID:          id,
		Description: description,
		CreatedAt:   timeProvider.GetTimeStamp(),
		IsComplete:  false,
	}
	vault.add(task)
	return id, nil
}

func GetTask(id uint, vault TaskVault) (Task, error) {
	task, err := vault.get(id)
	return task, err
}

func ListAllTasks(vault TaskVault) []Task {
	return vault.list()
}

func ListUnfinishedTasks(vault TaskVault) []Task {
	return vault.listUnfinished()
}
