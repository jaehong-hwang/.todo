package cli

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jaehong-hwang/todo/errors"
	"github.com/jaehong-hwang/todo/response"
	t "github.com/jaehong-hwang/todo/todo"
	"github.com/urfave/cli/v2"
)

func todoCommand(action cli.ActionFunc) cli.ActionFunc {
	return middleware(
		[]callback{
			todoRequireMiddleware,
			messageRequireMiddleware,
			authorSettingMiddleware,
		},
		action,
	)
}

func collectionCommand(action cli.ActionFunc) cli.ActionFunc {
	return middleware(
		[]callback{
			todoRequireMiddleware,
			authorSettingMiddleware,
		},
		action,
	)
}

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
			if c.NumFlags() == 0 {
				cli.ShowCommandHelpAndExit(c, "config", 0)
			}

			authorName := c.String("set-author-name")
			if authorName != "" {
				system.Author.Name = authorName
			}

			authorEmail := c.String("set-author-email")
			if authorEmail != "" {
				system.Author.Email = authorEmail
			}

			err := system.Save()
			if err == nil {
				appResponse = &response.MessageResponse{Message: "Your configuration has been saved successfully."}
			} else {
				appResponse = &response.ErrorResponse{Err: err}
			}
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
		Flags:   []cli.Flag{withDoneFlag, statusFlag, getJsonFlag},
		Aliases: []string{"l"},
		Usage:   "Print todos to the list",
		Action: collectionCommand(func(c *cli.Context) error {
			var col t.Collection

			status := c.String("status")

			if c.Bool("with-done") {
				col = *collection
			} else if status != "" {
				col = collection.SearchByStatus([]string{status})
			} else {
				col = collection.SearchByStatus([]string{t.StatusWaiting, t.StatusWorking})
			}

			if len(col.Todos) == 0 {
				return errors.New("todo_empty")
			}

			if c.Bool("get-json") {
				json, err := col.GetTodosJSONString()
				if err == nil {
					appResponse = &response.MessageResponse{Message: json}
				} else {
					appResponse = &response.ErrorResponse{Err: err}
				}
			} else {
				appResponse = &response.ListResponse{Todos: col.Todos}
			}
			return nil
		}),
	}

	addCommand = &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add todo",
		Action: todoCommand(func(c *cli.Context) error {
			todo := collection.NewTodo()
			todo.Content = c.Args().Get(0)
			todo.Status = t.StatusWaiting
			todo.Author = system.Author.Name
			todo.AuthorEmail = system.Author.Email
			todo.Start = time.Now()
			todo.End = time.Now()

			collection.Add(todo)

			return save()
		}),
	}

	addLabel = &cli.Command{
		Name:  "add-label",
		Flags: []cli.Flag{idFlag},
		Usage: "add label to todo",
		Action: todoCommand(func(c *cli.Context) error {
			id := c.Int("id")
			todo := collection.GetTodo(id)
			if todo == nil {
				return errors.NewWithParam("todo_id_not_found", map[string]string{
					"id": strconv.Itoa(id),
				})
			}

			labelText := c.Args().Get(0)
			label := t.Label{
				Text: labelText,
			}
			res := todo.AddLabel(&label)
			if res == false {
				return errors.NewWithParam("label_already_exists", map[string]string{
					"label": labelText,
				})
			}

			return save()
		}),
	}

	removeLabel = &cli.Command{
		Name:  "remove-label",
		Flags: []cli.Flag{idFlag},
		Usage: "remove label from todo",
		Action: todoCommand(func(c *cli.Context) error {
			id := c.Int("id")
			todo := collection.GetTodo(id)
			if todo == nil {
				return errors.NewWithParam("todo_id_not_found", map[string]string{
					"id": strconv.Itoa(id),
				})
			}

			labelText := c.Args().Get(0)
			res := todo.RemoveLabel(labelText)
			if res == false {
				return errors.NewWithParam("label_not_found", map[string]string{
					"label": labelText,
				})
			}

			return save()
		}),
	}

	updateCommand = &cli.Command{
		Name:    "update",
		Flags:   []cli.Flag{idFlag},
		Aliases: []string{"u"},
		Usage:   "update todo message",
		Action: todoCommand(func(c *cli.Context) error {
			id := c.Int("id")
			todo := collection.GetTodo(id)
			if todo == nil {
				return errors.NewWithParam("todo_id_not_found", map[string]string{
					"id": strconv.Itoa(id),
				})
			}

			todo.Content = c.Args().Get(0)

			return save()
		}),
	}

	removeCommand = &cli.Command{
		Name:    "remove",
		Flags:   []cli.Flag{idFlag},
		Aliases: []string{"rm"},
		Usage:   "remove todo message",
		Action: collectionCommand(func(c *cli.Context) error {
			id := c.Int("id")
			todo := collection.GetTodo(id)
			if todo == nil {
				return errors.NewWithParam("todo_id_not_found", map[string]string{
					"id": strconv.Itoa(id),
				})
			}

			yn := "y"
			fmt.Print("Do you want remove this todo?\nContent: ", todo.Content, " (y, n): ")
			fmt.Scanln(&yn)
			if yn != "y" && yn != "Y" {
				return nil
			}

			collection.Remove(id)

			return save()
		}),
	}

	removeCollectionCommand = &cli.Command{
		Name:  "remove-collection",
		Usage: "remove current todo collection",
		Action: collectionCommand(func(c *cli.Context) error {
			yn := "y"
			fmt.Print("Do you want remove current todo collection? (y, n): ")
			fmt.Scanln(&yn)
			if yn != "y" && yn != "Y" {
				return nil
			}

			return todoFile.Remove()
		}),
	}
)
