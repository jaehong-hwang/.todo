package cli

import (
	"errors"
	"os"
	"strconv"

	"github.com/jaehong-hwang/todo/response"
	t "github.com/jaehong-hwang/todo/todo"
	"github.com/urfave/cli/v2"
)

// TodoCommands are collection commands
type TodoCommands []*cli.Command

// GetCommands are making todo collection
func (a *App) GetCommands() TodoCommands {
	return []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "print todos to the list",
			Action:  a.list,
		},
		{
			Name:   "init",
			Usage:  "set up todo for current directory",
			Action: a.init,
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add todo",
			Action:  a.add,
		},
		{
			Name:    "state",
			Aliases: []string{"s"},
			Usage:   "update state",
			Action:  a.updateState,
		},
	}
}

func (a *App) list(c *cli.Context) error {
	a.response = &response.ListResponse{Todos: a.collection.Todos}
	return nil
}

func (a *App) init(c *cli.Context) error {
	if a.file.IsExists() {
		return errors.New("todo collection already exists")
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	err = a.file.CreateFile(dir)
	if err != nil {
		return err
	}

	a.response = &response.MessageResponse{Message: "todo init complete"}
	return nil
}

func (a *App) add(c *cli.Context) error {
	if !a.file.IsExists() {
		return errors.New("todo dosen't exists, you should run todo init")
	}

	todo := a.collection.NewTodo()
	todo.Content = c.Args().Get(0)
	todo.Status = t.StatusWaiting

	a.collection.Add(todo)

	content, err := a.collection.GetTodosByJSONString()
	if err != nil {
		return err
	}

	a.file.FillContent(content)

	return nil
}

func (a *App) updateState(c *cli.Context) error {
	id, err := strconv.Atoi(c.Args().Get(1))
	if err != nil {
		return err
	}

	todo := &a.collection.Todos[id]

	status := c.Args().Get(0)
	switch status {
	case "wait":
		todo.Status = t.StatusWaiting
	case "work":
		todo.Status = t.StatusWorking
	case "done":
		todo.Status = t.StatusDone
	default:
		return errors.New(status + " is unexpected state. todo have 3 state ex. wait, work, done")
	}

	content, err := a.collection.GetTodosByJSONString()
	if err != nil {
		return err
	}

	a.file.FillContent(content)

	return nil
}
