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
