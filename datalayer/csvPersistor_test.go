package datalayer

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestCsvRead(t *testing.T) {
	input := `ID,Description,CreatedAt,IsComplete
1,My new task,2024-07-27T16:45:19-05:00,true
2,Finish this video,2024-07-27T16:45:26-05:00,true
3,Find a video editor,2024-07-27T16:45:31-05:00,false`

	task1 := generateTask(1, "My new task", "2024-07-27T16:45:19-05:00", true)
	task2 := generateTask(2, "Finish this video", "2024-07-27T16:45:26-05:00", true)
	task3 := generateTask(3, "Find a video editor", "2024-07-27T16:45:31-05:00", false)

	want := MapTaskVault{
		db: map[uint]Task{
			task1.ID: task1,
			task2.ID: task2,
			task3.ID: task3,
		},
		lastId: task3.ID,
	}

	reader := strings.NewReader(input)
	got := CsvRead(reader)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %v, but got: %v", want, got)
	}
}

func generateTask(id uint, desctiption string, timeStr string, isComplete bool) Task {
	layout := time.RFC3339

	time, _ := time.Parse(layout, timeStr)
	return Task{
		ID:          id,
		Description: desctiption,
		CreatedAt:   time,
		IsComplete:  isComplete,
	}
}
