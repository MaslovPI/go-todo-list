package main

import (
	"log"
	"os"

	"github.com/maslovpi/go-todo-list/cmd"
	filemanagement "github.com/maslovpi/go-todo-list/fileManagement"
	"github.com/maslovpi/go-todo-list/logging"
)

var logger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)

func main() {
	path, err := filemanagement.GetCSVPath()
	logging.LogError(err)

	file, err := filemanagement.LoadFile(path)
	logging.LogError(err)

	defer filemanagement.CloseFile(file)

	cmd.Execute()
}
