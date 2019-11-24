package main

import (
	"errors"
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
	commands   *cli.App
	collection *TodoCollection
}

// NewApp find file and returns app
func NewApp() *App {
	file := &File{name: ".todo", permission: 0644}
	if err := file.FindFromCurrentDirectory(); err != nil {
		panic(err)
	}

	collection := NewTodoCollection(file)

	commands := &cli.App{
		Name:      "todo",
		Copyright: "(c) 2019 JaeHong Hwang",
		HelpName:  "contrive",
		Usage:     "",
		UsageText: `Todo app helper, You can run the following commands.`,
		Version:   "0.0.1",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "print todos to the list",
				Action: func(c *cli.Context) error {
					ResponseChan <- &ListResponse{todos: collection.Todos}
					return nil
				},
			},
			{
				Name:  "init",
				Usage: "set up todo for current directory",
				Action: func(c *cli.Context) error {
					if file.IsExists() {
						return errors.New("todo collection already exists")
					}

					dir, err := os.Getwd()
					if err != nil {
						return err
					}

					err = file.CreateFile(dir)
					if err != nil {
						return err
					}

					ResponseChan <- &MessageResponse{message: "todo init complete"}
					return nil
				},
			},
		},
	}

	return &App{
		collection: collection,
		commands:   commands,
	}
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

	if err := a.commands.Run(args); err != nil {
		ResponseChan <- &ErrorResponse{err: err}
	}
}
