package cli

import (
	"fmt"
	"os"

	"github.com/jaehong-hwang/todo/errors"
	"github.com/jaehong-hwang/todo/response"
	t "github.com/jaehong-hwang/todo/todo"
	"github.com/urfave/cli/v2"
)

// TodoCommands are collection commands
type TodoCommands []*cli.Command

var (
	initCommand = &cli.Command{
		Name:  "init",
		Usage: "set up todo for current directory",
		Action: func(c *cli.Context) error {
			if todoFile.IsExist() {
				return errors.New("todo_already_exists")
			}

			dir, err := os.Getwd()
			if err != nil {
				return err
			}

			err = todoFile.CreateIfNotExist()
			if err != nil {
				return err
			}

			system.AddDirectory(dir)

			appResponse = &response.MessageResponse{Message: "todo init complete"}
			return nil
		},
	}

	configCommand = &cli.Command{
		Name:    "config",
		Flags:   []cli.Flag{setAuthorName, setAuthorEmail},
		Aliases: []string{"c"},
		Usage:   "Set config global todo",
		Action: func(c *cli.Context) error {
			authorName := c.String("set-author-name")
			if authorName != "" {
				system.Author.Name = authorName
			}

			authorEmail := c.String("set-author-email")
			if authorEmail != "" {
				system.Author.Email = authorEmail
			}

			system.Save()

			return nil
		},
	}

	directoriesCommand = &cli.Command{
		Name:    "directories",
		Aliases: []string{"d"},
		Usage:   "Print directories of todo collection",
		Action: func(c *cli.Context) error {
			appResponse = &response.DirectoryResponse{Directories: system.Directories}
			return nil
		},
	}

	listCommand = &cli.Command{
		Name:    "list",
		Flags:   []cli.Flag{withDoneFlag, statusFlag},
		Aliases: []string{"l"},
		Usage:   "Print todos to the list",
		Action: func(c *cli.Context) error {
			var todos t.Todos

			status := c.String("status")

			if c.Bool("with-done") {
				todos = collection.Todos
			} else if status != "" {
				todos = collection.GetTodosByStatus([]string{status})
			} else {
				todos = collection.GetTodosByStatus([]string{t.StatusWaiting, t.StatusWorking})
			}

			appResponse = &response.ListResponse{Todos: todos}
			return nil
		},
	}

	addCommand = &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add todo",
		Action: func(c *cli.Context) error {
			if todoFile == nil {
				return errors.New("todo_doesnt_exists")
			}

			if c.NArg() == 0 {
				return errors.New("message_required")
			}

			if system.Author.Name == "" || system.Author.Email == "" {
				fmt.Println("[Warning] You should set up author setting, `todo config`")
			}

			todo := collection.NewTodo()
			todo.Content = c.Args().Get(0)
			todo.Status = t.StatusWaiting
			todo.Author = system.Author.Name
			todo.AuthorEmail = system.Author.Email

			collection.Add(todo)

			content, err := collection.GetTodosByJSONString()
			if err != nil {
				return err
			}

			return todoFile.FillContent(content)
		},
	}

	updateCommand = &cli.Command{
		Name:    "update",
		Flags:   []cli.Flag{idFlag},
		Aliases: []string{"u"},
		Usage:   "update todo message",
		Action: func(c *cli.Context) error {
			if todoFile == nil {
				return errors.New("todo_doesnt_exists")
			}

			if c.NArg() == 0 {
				return errors.New("message_required")
			}

			id := c.Int("id")
			todo := &collection.Todos[id]
			todo.Content = c.Args().Get(0)

			content, err := collection.GetTodosByJSONString()
			if err != nil {
				return err
			}

			return todoFile.FillContent(content)
		},
	}

	removeCommand = &cli.Command{
		Name:    "remove",
		Flags:   []cli.Flag{idFlag},
		Aliases: []string{"rm"},
		Usage:   "remove todo message",
		Action: func(c *cli.Context) error {
			if todoFile == nil {
				return errors.New("todo_doesnt_exists")
			}

			id := c.Int("id")
			todo := collection.Todos[id]
			yn := "y"
			fmt.Print("Do you want remove this todo?\n Content: ", todo.Content, " (y, n): ")
			fmt.Scanln(&yn)
			if yn != "y" && yn != "Y" {
				return nil
			}

			collection.Remove(id)

			content, err := collection.GetTodosByJSONString()
			if err != nil {
				return err
			}

			return todoFile.FillContent(content)
		},
	}
)
