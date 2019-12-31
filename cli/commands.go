package cli

import (
	"errors"
	"os"
	"strconv"

	"github.com/jaehong-hwang/todo/file"
	"github.com/jaehong-hwang/todo/response"
	t "github.com/jaehong-hwang/todo/todo"
	"github.com/urfave/cli/v2"
)

// TodoCommands are collection commands
type TodoCommands []*cli.Command

var (
	listCommand = &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "Print todos to the list",
		Action: func(c *cli.Context) error {
			appResponse = &response.ListResponse{Todos: collection.Todos}
			return nil
		},
	}
)

// GetCommands are making todo collection
func (a *App) GetCommands() TodoCommands {
	return []*cli.Command{
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
			Action: func(c *cli.Context) error {
				id, err := strconv.Atoi(c.Args().Get(1))
				if err != nil {
					return err
				}

				status := c.Args().Get(0)

				return a.updateState(id, status)
			},
		},
		{
			Name:  "wait",
			Usage: "todo set waiting state",
			Action: func(c *cli.Context) error {
				id, err := strconv.Atoi(c.Args().First())
				if err != nil {
					return err
				}

				return a.updateState(id, "wait")
			},
		},
		{
			Name:  "work",
			Usage: "todo set working state",
			Action: func(c *cli.Context) error {
				id, err := strconv.Atoi(c.Args().First())
				if err != nil {
					return err
				}

				return a.updateState(id, "work")
			},
		},
		{
			Name:  "done",
			Usage: "todo set done state",
			Action: func(c *cli.Context) error {
				id, err := strconv.Atoi(c.Args().First())
				if err != nil {
					return err
				}

				return a.updateState(id, "done")
			},
		},
	}
}

func (a *App) init(c *cli.Context) error {
	if a.file.IsExists() {
		return errors.New("todo collection already exists")
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	err = file.CreateTodoFile(dir)
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

func (a *App) updateState(id int, status string) error {
	todo := &a.collection.Todos[id]

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
