package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/urfave/cli/v2"
)

const (
	// TodoFileName is name of todo collection file
	todoFileName string = ".todo"

	// TodoFilePermission set read permission
	todoFilePermission os.FileMode = 0644

	// TodoNotFound error message
	todoNotFound string = "todo collection doesn't exists, please run 'todo init'"
)

// App is command center
type App struct {
	cliApp     *cli.App
	collection *TodoCollection
	file       *File
}

// NewApp find file and returns app
func NewApp() *App {
	app := App{}

	app.file = &File{name: ".todo", permission: 0644}
	if err := app.file.FindFromCurrentDirectory(); err != nil {
		panic(err)
	}

	app.collection = NewTodoCollection(app.file)

	app.cliApp = &cli.App{
		Name:      "todo",
		Copyright: "(c) 2019 JaeHong Hwang",
		HelpName:  "contrive",
		Usage:     "",
		UsageText: `Todo app helper, You can run the following commands.`,
		Version:   "0.0.1",
		Commands:  app.GetCommands(),
	}

	return &app
}

// Run to running correct command
func (a *App) Run(args []string, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()

		if err := recover(); err != nil {
			fmt.Println(err)
			runtime.Goexit()
		}
	}()

	if err := a.cliApp.Run(args); err != nil {
		ResponseChan <- &ErrorResponse{err: err}
	}
}
