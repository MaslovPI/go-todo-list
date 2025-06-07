package datalayer

import (
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
