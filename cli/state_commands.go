package cli

import (
	"errors"
	"strconv"

	t "github.com/jaehong-hwang/todo/todo"
	"github.com/urfave/cli/v2"
)

var (
	stateCommand = &cli.Command{
		Name:    "state",
		Aliases: []string{"s"},
		Usage:   "update state",
		Action: func(c *cli.Context) error {
			id, err := strconv.Atoi(c.Args().Get(1))
			if err != nil {
				return err
			}

			status := c.Args().Get(0)

			return updateState(id, status)
		},
	}

	waitCommand = &cli.Command{
		Name:   "wait",
		Usage:  "todo set waiting state",
		Action: getUpdatingStateAction("wait"),
	}

	workCommand = &cli.Command{
		Name:   "work",
		Usage:  "todo set working state",
		Action: getUpdatingStateAction("work"),
	}

	doneCommand = &cli.Command{
		Name:   "done",
		Usage:  "todo set done state",
		Action: getUpdatingStateAction("done"),
	}
)

func getUpdatingStateAction(state string) func(*cli.Context) error {
	return func(c *cli.Context) error {
		id, err := strconv.Atoi(c.Args().First())
		if err != nil {
			return err
		}

		return updateState(id, state)
	}
}

func updateState(id int, status string) error {
	todo := &collection.Todos[id]

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

	content, err := collection.GetTodosByJSONString()
	if err != nil {
		return err
	}

	todoFile.FillContent(content)

	return nil
}
