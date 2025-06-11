package datalayer

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

const (
	ErrReaderError = MockError("Error with reader")
)

type MockError string

func (e MockError) Error() string {
	return string(e)
}

type FaultyReader struct{}

func (f FaultyReader) Read(p []byte) (n int, err error) {
	return 0, ErrReaderError
}

func TestCsvRead(t *testing.T) {
	t.Run("Read correct csv", func(t *testing.T) {
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
		got, err := CsvRead(reader)
		assertNoError(t, err)
		assertMapVaultsEqual(t, want, got)
	})

	t.Run("Can't read csv", func(t *testing.T) {
		reader := FaultyReader{}
		_, err := CsvRead(reader)
		assertError(t, err, ErrReaderError)
	})

	t.Run("read with incorrect id", func(t *testing.T) {
		input := `ID,Description,CreatedAt,IsComplete
NAN,My new task,2024-07-27T16:45:19-05:00,true`

		reader := strings.NewReader(input)
		_, err := CsvRead(reader)

		assertNumError(t, err)
	})

	t.Run("read with incorrect date", func(t *testing.T) {
		input := `ID,Description,CreatedAt,IsComplete
1,My new task,notADate,true`

		reader := strings.NewReader(input)
		_, err := CsvRead(reader)

		assertParseError(t, err)
	})

	t.Run("read with incorrect bool", func(t *testing.T) {
		input := `ID,Description,CreatedAt,IsComplete
1,My new task,2024-07-27T16:45:19-05:00,notABool`

		reader := strings.NewReader(input)
		_, err := CsvRead(reader)

		assertNumError(t, err)
	})
}

func TestCsvWrite(t *testing.T) {
	t.Run("Write correct csv", func(t *testing.T) {
		task1 := generateTask(1, "My new task", "2024-07-27T16:45:19-05:00", true)
		task2 := generateTask(2, "Finish this video", "2024-07-27T16:45:26-05:00", true)
		task3 := generateTask(3, "Find a video editor", "2024-07-27T16:45:31-05:00", false)
		input := MapTaskVault{
			db: map[uint]Task{
				task1.ID: task1,
				task2.ID: task2,
				task3.ID: task3,
			},
			lastId: task3.ID,
		}

		buf := new(bytes.Buffer)
		err := CsvWrite(input, buf)
		assertNoError(t, err)

		got := buf.String()
		want := `ID,Description,CreatedAt,IsComplete
1,My new task,2024-07-27T16:45:19-05:00,true
2,Finish this video,2024-07-27T16:45:26-05:00,true
3,Find a video editor,2024-07-27T16:45:31-05:00,false
`
		assertStringEqual(t, got, want)
	})
}

func generateTask(id uint, desctiption string, timeStr string, isComplete bool) Task {
	time, _ := time.Parse(layout, timeStr)
	return Task{
		ID:          id,
		Description: desctiption,
		CreatedAt:   time,
		IsComplete:  isComplete,
	}
}
