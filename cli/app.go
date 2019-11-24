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

// NewApp find file and returns app
func NewApp() *App {
	app := App{}

	app.file = &file.File{Name: ".todo", Permission: 0644}
	if err := app.file.FindFromCurrentDirectory(); err != nil {
		panic(err)
	}

	app.collection = todo.NewTodoCollection(app.file)

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
func (a *App) Run(args []string) response.Response {
	if err := a.cliApp.Run(args); err != nil {
		return &response.ErrorResponse{Err: err}
	}

	return a.response
}
