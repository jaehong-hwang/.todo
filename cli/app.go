package cli

import (
	"github.com/jaehong-hwang/todo/file"
	"github.com/jaehong-hwang/todo/response"
	"github.com/jaehong-hwang/todo/todo"
	"github.com/urfave/cli/v2"
)

// App is command center
type App struct {
	response   response.Response
	cliApp     *cli.App
	collection *todo.Collection
	file       *file.File
}

var (
	app        = newApp()
	todoFile   = file.FindTodoFile()
	todoSystemFile = file.FindTodoSystemFile()
	collection = todo.NewTodoCollection(todoFile)
)

var appResponse response.Response

// newApp find file and returns app
func newApp() *cli.App {
	cliApp := &cli.App{
		Name:      "todo",
		Copyright: "(c) 2019 JaeHong Hwang",
		HelpName:  "contrive",
		Usage:     "",
		UsageText: `Todo app helper, You can run the following commands.`,
		Version:   "0.0.1",
		Commands: []*cli.Command{
			initCommand,
			listCommand,
			addCommand,
			updateCommand,
			stateCommand,
			waitCommand,
			workCommand,
			doneCommand,
		},
	}

	return cliApp
}

// Run to running correct command
func Run(args []string) response.Response {
	if err := app.Run(args); err != nil {
		return &response.ErrorResponse{Err: err}
	}

	return appResponse
}
