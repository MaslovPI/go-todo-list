package datalayer

import (
	"reflect"
	"testing"
	"time"
)

type DummyTimeProvider struct {
	timeStamp time.Time
}

func (d *DummyTimeProvider) GetTimeStamp() time.Time {
	return d.timeStamp
}

func TestAdd(t *testing.T) {
	vault := NewMapTaskVault()
	timeProvider := DummyTimeProvider{}
	timeProvider.timeStamp = time.Now()

	t.Run("New task with description", func(t *testing.T) {
		description := "test"
		id, err := AddTask(description, vault, &timeProvider)
		assertNoError(t, err)
		got, err := GetTask(id, vault)
		assertNoError(t, err)

		want := Task{
			ID:          1,
			Description: description,
			CreatedAt:   timeProvider.timeStamp,
			IsComplete:  false,
		}
		assertTasksEqual(t, got, want)
	})

	t.Run("New task without description", func(t *testing.T) {
		description := ""
		_, err := AddTask(description, vault, &timeProvider)
		assertError(t, err, ErrTaskDescriptionEmpty)
	})
}

func TestGet(t *testing.T) {
	timeProvider := DummyTimeProvider{}
	timeProvider.timeStamp = time.Now()
	t.Run("Task exists", func(t *testing.T) {
		var id uint = 999
		want := Task{
			ID:          id,
			Description: "test",
			CreatedAt:   timeProvider.timeStamp,
			IsComplete:  false,
		}

		vault := MapTaskVault{
			db:     map[uint]Task{id: want},
			lastId: id,
		}

		got, err := GetTask(id, &vault)
		assertNoError(t, err)
		assertTasksEqual(t, got, want)
	})
	t.Run("Task doesn't exists", func(t *testing.T) {
		var id uint = 999
		vault := NewMapTaskVault()

		_, err := GetTask(id, vault)
		assertError(t, err, ErrNotFound)
	})
}

func TestListAllTasks(t *testing.T) {
	timeProvider := DummyTimeProvider{}
	timeProvider.timeStamp = time.Now()
	t.Run("List tasks full vault", func(t *testing.T) {
		taskIncomplete := Task{
			ID:          998,
			Description: "test",
			CreatedAt:   timeProvider.timeStamp,
			IsComplete:  false,
		}
		taskComplete := Task{
			ID:          999,
			Description: "test complete",
			CreatedAt:   timeProvider.timeStamp,
			IsComplete:  true,
		}

		vault := MapTaskVault{
			db: map[uint]Task{
				taskIncomplete.ID: taskIncomplete,
				taskComplete.ID:   taskComplete,
			},
			lastId: taskComplete.ID,
		}

		got := ListAllTasks(&vault)
		assertNumberEqual(t, len(got), 2)
		assertTasksEqual(t, got[0], taskIncomplete)
		assertTasksEqual(t, got[1], taskComplete)
	})
	t.Run("List tasks empty vault", func(t *testing.T) {
		vault := NewMapTaskVault()

		got := ListAllTasks(vault)
		assertNumberEqual(t, len(got), 0)
	})
}

func TestListUnfinishedTasks(t *testing.T) {
	timeProvider := DummyTimeProvider{}
	timeProvider.timeStamp = time.Now()
	t.Run("List unfinished tasks full vault", func(t *testing.T) {
		taskIncomplete := Task{
			ID:          998,
			Description: "test",
			CreatedAt:   timeProvider.timeStamp,
			IsComplete:  false,
		}
		taskComplete := Task{
			ID:          999,
			Description: "test complete",
			CreatedAt:   timeProvider.timeStamp,
			IsComplete:  true,
		}

		vault := MapTaskVault{
			db: map[uint]Task{
				taskIncomplete.ID: taskIncomplete,
				taskComplete.ID:   taskComplete,
			},
			lastId: taskComplete.ID,
		}

		got := ListUnfinishedTasks(&vault)
		assertNumberEqual(t, len(got), 1)
		assertTasksEqual(t, got[0], taskIncomplete)
	})
	t.Run("List unfinished tasks empty vault", func(t *testing.T) {
		vault := NewMapTaskVault()

		got := ListUnfinishedTasks(vault)
		assertNumberEqual(t, len(got), 0)
	})
}

func TestDelete(t *testing.T) {
	timeProvider := DummyTimeProvider{}
	timeProvider.timeStamp = time.Now()
	t.Run("Task exists", func(t *testing.T) {
		var id uint = 999
		want := Task{
			ID:          id,
			Description: "test",
			CreatedAt:   timeProvider.timeStamp,
			IsComplete:  false,
		}

		vault := MapTaskVault{
			db:     map[uint]Task{id: want},
			lastId: id,
		}
		err := DeleteTask(id, &vault)
		assertNoError(t, err)

		_, err = GetTask(id, &vault)
		assertError(t, err, ErrNotFound)
	})
	t.Run("Task doesn't exists", func(t *testing.T) {
		var id uint = 999
		vault := NewMapTaskVault()

		err := DeleteTask(id, vault)
		assertError(t, err, ErrTaskDoesNotExist)
	})
}

func TestComplete(t *testing.T) {
	timeProvider := DummyTimeProvider{}
	timeProvider.timeStamp = time.Now()
	t.Run("Uncomplete task exists", func(t *testing.T) {
		var id uint = 999
		originalTask := Task{
			ID:          id,
			Description: "test",
			CreatedAt:   timeProvider.timeStamp,
			IsComplete:  false,
		}

		vault := MapTaskVault{
			db:     map[uint]Task{id: originalTask},
			lastId: id,
		}
		err := CompleteTask(id, &vault)
		assertNoError(t, err)

		got, err := GetTask(id, &vault)
		assertNoError(t, err)
		want := originalTask
		want.IsComplete = true

		assertTasksEqual(t, got, want)
	})
	t.Run("Complete task exists", func(t *testing.T) {
		var id uint = 999
		originalTask := Task{
			ID:          id,
			Description: "test",
			CreatedAt:   timeProvider.timeStamp,
			IsComplete:  true,
		}

		vault := MapTaskVault{
			db:     map[uint]Task{id: originalTask},
			lastId: id,
		}
		err := CompleteTask(id, &vault)
		assertError(t, err, ErrTaskAlreadyComplete)
	})
	t.Run("Task doesn't exists", func(t *testing.T) {
		var id uint = 999
		vault := NewMapTaskVault()

		err := CompleteTask(id, vault)
		assertError(t, err, ErrTaskDoesNotExist)
	})
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Didn't expect error, but got: %v", err)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Fatalf("got error %q want %q", got, want)
	}
}

func assertTasksEqual(t testing.TB, got, want Task) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Want: %#v, but got: %#v", want, got)
	}
}

func assertNumberEqual(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("Want: %d, but got: %d", want, got)
	}
}
