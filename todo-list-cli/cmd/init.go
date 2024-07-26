package cmd

import (
	"todo-list-cli/data"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		data.CreateTableCategories()
		data.CreateTableNotes()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
