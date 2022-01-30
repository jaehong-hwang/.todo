package cli

import "github.com/urfave/cli/v2"

var (
	idFlag = &cli.IntFlag{
		Name:     "id",
		Aliases:  []string{"i"},
		Usage:    "Select todo by ID",
		Required: true,
	}

	withDoneFlag = &cli.BoolFlag{
		Name:  "with-done",
		Usage: "Do you want already completed task?",
	}

	statusFlag = &cli.StringFlag{
		Name:    "status",
		Aliases: []string{"s"},
		Usage:   "Search by inputed status",
	}

	setAuthorName = &cli.StringFlag{
		Name:  "set-author-name",
		Usage: "literally set global author name",
	}

	setAuthorEmail = &cli.StringFlag{
		Name:  "set-author-email",
		Usage: "literally set global author email",
	}

	getJsonFlag = &cli.BoolFlag{
		Name:  "get-json",
		Usage: "return in Json format",
	}

	dirFlag = &cli.StringFlag{
		Name:  "directory",
		Usage: "action with directory flag",
	}
)
