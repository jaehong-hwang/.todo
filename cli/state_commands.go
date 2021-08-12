package cli

import (
	"strconv"

	"github.com/jaehong-hwang/todo/errors"
	t "github.com/jaehong-hwang/todo/todo"
	"github.com/urfave/cli/v2"
)

var (
	stateCommand = &cli.Command{
		Name:    "state",
		Aliases: []string{"s"},
		Flags:   []cli.Flag{idFlag},
		Usage:   "update state",
		Action: func(c *cli.Context) error {
			id := c.Int("id")
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
	todo := collection.GetTodo(id)
	if todo == nil {
		return errors.NewWithParam("todo_id_not_found", map[string]string{
			"id": strconv.Itoa(id),
		})
	}

	switch status {
	case "wait":
		todo.Status = t.StatusWaiting
	case "work":
		todo.Status = t.StatusWorking
	case "done":
		todo.Status = t.StatusDone
	default:
		return errors.NewWithParam("unexpected_state", map[string]string{
			"state": status,
		})
	}

	content, err := collection.GetTodosJSONString()
	if err != nil {
		return err
	}

	todoFile.FillContent(content)

	return nil
}
