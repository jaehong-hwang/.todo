package cli

import "github.com/urfave/cli/v2"

var (
  idFlag = &cli.IntFlag{
	Name: "id",
	Aliases: []string{"i"},
	Usage: "Select todo by ID",
	Required: true,
  }
)
