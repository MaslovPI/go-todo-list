package servicelayer

import (
	"fmt"
	"os"

	"github.com/maslovpi/go-todo-list/datalayer"
	filemanagement "github.com/maslovpi/go-todo-list/fileManagement"
	"github.com/maslovpi/go-todo-list/logging"
)

type ToDoList struct {
	vault datalayer.MapTaskVault
	file  *os.File
}

func NewToDoList() *ToDoList {
	path, err := filemanagement.GetCSVPath()
	logging.LogFatal(err)
	file, err := filemanagement.LoadFile(path)
	logging.LogFatal(err)
	mapVault, err := datalayer.CsvRead(file)
	logging.LogFatal(err)
	return &ToDoList{vault: mapVault}
}

func (t *ToDoList) Finalize() {
	err := datalayer.CsvWrite(t.vault, t.file)
	logging.LogFatal(err)
	err = filemanagement.CloseFile(t.file)
	logging.LogFatal(err)
}

func (t *ToDoList) Add(description string) {
	id, err := datalayer.AddTask(description, &t.vault, &datalayer.DefaultTimeProvider{})
	logging.LogError(err)
	logging.LogInfo(fmt.Sprintf("Task added (ID =%d, Description = %q)", id, description))
}

func (t *ToDoList) ListAll() {
	datalayer.ListAllTasks(&t.vault)
}

func (t *ToDoList) ListUnfinished() {
	datalayer.ListUnfinishedTasks(&t.vault)
}

func (t *ToDoList) Complete(id uint) {
	err := datalayer.CompleteTask(id, &t.vault)
	logging.LogError(err)
}

func (t *ToDoList) Delete(id uint) {
	err := datalayer.DeleteTask(id, &t.vault)
	logging.LogError(err)
}
