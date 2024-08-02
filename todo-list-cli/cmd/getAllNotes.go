/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"todo-list-cli/data"
)

// getAllNotesCmd represents the getAllNotes command
var getAllNotesCmd = &cobra.Command{
	Use:   "getAllNotes",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		notes, err := data.FindAllNotes()
		if err != nil {
			fmt.Println(err)
			return
		}
		for i, note := range notes {
			fmt.Printf("%v. Note: %v\n   Description: %v\n   Category: %v\n", i+1, note.NoteName, note.Description, note.CategoryName)
		}
	},
}

func init() {
	rootCmd.AddCommand(getAllNotesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getAllNotesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getAllNotesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
