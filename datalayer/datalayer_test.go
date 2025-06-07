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
	vault := MapTaskVault{}
	vault.init()
	timeProvider := DummyTimeProvider{}
	timeProvider.timeStamp = time.Now()

	description := "test"
	id := AddTask(description, &vault, &timeProvider)

	got, err := GetTask(id, &vault)
	assertNoError(t, err)

	want := Task{
		ID:          1,
		Description: description,
		CreatedAt:   timeProvider.timeStamp,
		IsComplete:  false,
	}
	assertTasksEqual(t, got, want)
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
		vault := MapTaskVault{}
		vault.init()

		_, err := GetTask(id, &vault)
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
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertTasksEqual(t testing.TB, got, want Task) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Want: %#v, but got: %#v", want, got)
	}
}
