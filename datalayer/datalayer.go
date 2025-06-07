package datalayer

import (
	"time"
)

const (
	ErrNotFound             = VaultErr("Could not find task")
	ErrTaskDoesNotExist     = VaultErr("Cannot perform operation on task because it does not exist")
	ErrTaskDescriptionEmpty = VaultErr("Task description cannot be empty")
	ErrTaskAlreadyComplete  = VaultErr("Cannot complete task. Already complete.")
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
	vault.addOrUpdate(task)
	return id, nil
}

func GetTask(id uint, vault TaskVault) (Task, error) {
	task, ok := vault.get(id)
	if !ok {
		return Task{}, ErrNotFound
	}
	return task, nil
}

func ListAllTasks(vault TaskVault) []Task {
	return vault.list()
}

func ListUnfinishedTasks(vault TaskVault) []Task {
	return vault.listUnfinished()
}

func DeleteTask(id uint, vault TaskVault) error {
	if !vault.exists(id) {
		return ErrTaskDoesNotExist
	}
	vault.delete(id)
	return nil
}

func CompleteTask(id uint, vault TaskVault) error {
	task, ok := vault.get(id)
	if !ok {
		return ErrTaskDoesNotExist
	}
	if task.IsComplete {
		return ErrTaskAlreadyComplete
	}

	task.IsComplete = true
	vault.addOrUpdate(task)
	return nil
}
