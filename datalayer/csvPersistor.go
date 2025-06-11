package datalayer

import (
	"encoding/csv"
	"io"
	"strconv"
	"time"
)

func CsvRead(reader io.Reader) (MapTaskVault, error) {
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = 4
	csvData, err := csvReader.ReadAll()
	if err != nil {
		return MapTaskVault{}, err
	}

	dictionary := make(map[uint]Task)
	var maxId uint = 0
	for i, each := range csvData {
		if i == 0 {
			continue
		}

		task, err := stringArrayToTask(each)
		if err != nil {
			return MapTaskVault{}, err
		}

		dictionary[task.ID] = task

		if task.ID > maxId {
			maxId = task.ID
		}
	}

	return MapTaskVault{db: dictionary, lastId: maxId}, nil
}

func CsvWrite(vault MapTaskVault, writer io.Writer) {
	csvWriter := csv.NewWriter(writer)
	csvWriter.Write(getTitleRow())
	for _, task := range vault.list() {
		record := taskToStringArray(task)
		csvWriter.Write(record)
	}
	csvWriter.Flush()
}

func getTitleRow() []string {
	result := make([]string, 4)
	result[0] = "ID"
	result[1] = "Description"
	result[2] = "CreatedAt"
	result[3] = "IsComplete"
	return result
}

func taskToStringArray(task Task) []string {
	result := make([]string, 4)
	result[0] = strconv.Itoa(int(task.ID))
	result[1] = task.Description
	result[2] = task.CreatedAt.Format(layout)
	result[3] = strconv.FormatBool(task.IsComplete)
	return result
}

const layout = time.RFC3339

func stringArrayToTask(array []string) (Task, error) {
	id, err := strconv.Atoi(array[0])
	if err != nil {
		return Task{}, err
	}

	description := array[1]
	time, err := time.Parse(layout, array[2])
	if err != nil {
		return Task{}, err
	}

	isComplete, err := strconv.ParseBool(array[3])
	if err != nil {
		return Task{}, err
	}

	return Task{
		ID:          uint(id),
		Description: description,
		CreatedAt:   time,
		IsComplete:  isComplete,
	}, nil
}
