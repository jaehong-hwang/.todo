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
	app            = newApp()
	todoFile       = file.FindTodoFile()
	todoSystemFile = file.FindTodoSystemFile()
	collection     = todo.NewTodoCollection(todoFile)
	system         = todo.NewSystem(todoSystemFile)
	isJson         bool
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
			configCommand,
			directoriesCommand,
			listCommand,
			addCommand,
			updateCommand,
			addLabel,
			removeLabel,
			stateCommand,
			waitCommand,
			workCommand,
			doneCommand,
			removeCommand,
			removeCollectionCommand,
		},
		After: func(c *cli.Context) error {
			for _, val := range c.Args().Slice() {
				if val == "--get-json" {
					isJson = true
				}
			}
			return nil
		},
	}

	return cliApp
}

// Run to running correct command, first is
func Run(args []string) (response.Response, bool) {
	if err := app.Run(args); err != nil {
		return &response.ErrorResponse{Err: err}, isJson
	}

	return appResponse, isJson
}
