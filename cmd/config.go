package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Set config global todo",
	}

	configNameCmd = &cobra.Command{
		Use:   "name",
		Short: "Set author name",
		Args:  requireArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			system.Author.Name = args[0]
			return saveSystem()
		},
	}

	configEmailCmd = &cobra.Command{
		Use:   "email",
		Short: "Set author email",
		Args:  requireArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			system.Author.Email = args[0]
			return saveSystem()
		},
	}
)

func saveSystem() error {
	err := system.Save()
	if err != nil {
		return err
	}

	fmt.Println("Your configuration has been saved successfully.")
	return nil
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configNameCmd)
	configCmd.AddCommand(configEmailCmd)
}
