package datalayer

import (
	"cmp"
	"maps"
	"slices"
)

type TaskVault interface {
	addOrUpdate(task Task)
	get(id uint) (Task, bool)
	getNextId() uint
	list() []Task
	listUnfinished() []Task
	delete(id uint)
	exists(id uint) bool
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

func (m *MapTaskVault) addOrUpdate(task Task) {
	m.db[task.ID] = task
}

func (m *MapTaskVault) get(id uint) (Task, bool) {
	task, ok := m.db[id]
	return task, ok
}

func (m *MapTaskVault) list() []Task {
	slice := slices.Collect(maps.Values(m.db))
	slices.SortFunc(slice, func(a, b Task) int {
		return cmp.Compare(a.ID, b.ID)
	})
	return slice
}

func (m *MapTaskVault) listUnfinished() []Task {
	allTasks := m.list()
	result := make([]Task, 0, len(m.db))
	for _, task := range allTasks {
		if !task.IsComplete {
			result = append(result, task)
		}
	}
	return result
}

func (m *MapTaskVault) delete(id uint) {
	delete(m.db, id)
}

func (m *MapTaskVault) exists(id uint) bool {
	_, ok := m.db[id]
	return ok
}

func (m *MapTaskVault) getNextId() uint {
	m.lastId++
	return m.lastId
}
