package main

import (
	"log"
	"os"

	"github.com/maslovpi/go-todo-list/cmd"
)

var logger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)

func main() {
	cmd.Execute()
}
