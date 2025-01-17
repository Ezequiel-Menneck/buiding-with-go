/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"todo-list-cli/data"
	"todo-list-cli/promptui_utils"
)

var noteName string
var description string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(noteName) == 0 || len(description) == 0 {
			fmt.Println("Please enter a note name/description")
			return
		}
		note, err := data.GetNoteByName(noteName)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			fmt.Println("something went wrong trying to get note by name")
			return
		}

		if note != nil {
			fmt.Println("note already exists")
			return
		}

		categoryName, err := promptui_utils.PrintPromptUiCategories()

		id, err := data.FindCategoryByName(*categoryName)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			fmt.Println("something went wrong trying to get category")
			return
		}

		if err != nil && errors.Is(err, sql.ErrNoRows) {
			id, err = data.CreateCategory(*categoryName)
			if err != nil {
				fmt.Println("something went wrong trying to create category")
				return
			}
		}

		note, err = data.InsertNote(data.Note{
			NoteName:    noteName,
			Description: description,
			CategoryId:  id,
		})
		if err != nil {
			return
		}

		fmt.Printf("Note added successfully: \n\tNote: %v\n\tDescription: %v\n\tCategory: %v\n", note.NoteName, note.Description, categoryName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&noteName, "note", "n", "", "note to add")
	addCmd.Flags().StringVarP(&description, "description", "d", "", "description of your note")
}
