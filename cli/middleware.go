package cli

import (
	"fmt"

	"github.com/jaehong-hwang/todo/errors"
	"github.com/jaehong-hwang/todo/file"
	"github.com/jaehong-hwang/todo/todo"
	"github.com/urfave/cli/v2"
)

type callback func(*cli.Context) error

func errorCommandAction(err error) cli.ActionFunc {
	return func(c *cli.Context) error {
		return err
	}
}

func middleware(middlewares []callback, action cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		for _, middlewareFunc := range middlewares {
			err := middlewareFunc(c)
			if err != nil {
				return err
			}
		}

		return action(c)
	}
}

func todoRequireMiddleware(c *cli.Context) error {
	if todoFile.IsExist() == false {
		return errors.New("todo_doesnt_exists")
	}

	return nil
}

func messageRequireMiddleware(c *cli.Context) error {
	if c.NArg() == 0 {
		return errors.New("message_required")
	}

	return nil
}

func authorSettingMiddleware(c *cli.Context) error {
	if system.Author.Name == "" || system.Author.Email == "" {
		fmt.Println("[Warning] You should set up author setting, `todo config`")
	}

	return nil
}

func userDirectoryMiddleware(c *cli.Context) error {
	dir := c.String("directory")
	todoFile = file.FindTodoFileWithDirectory(dir)
	collection = todo.NewTodoCollection(todoFile)

	return nil
}
