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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		note, err := promptui_utils.PrintPrompUiNotes()
		if err != nil {
			fmt.Printf("Prompt UI note could not be displayed, error %v\n", err)
			return
		}

		noteName := promptui_utils.UserInputToUpdateNote("Name")

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

		description := promptui_utils.UserInputToUpdateNote("Description")

		_, err = data.UpdateNote(data.NoteUpdate{
			NoteName:    *note,
			NewNoteName: noteName,
			Description: description,
			CategoryId:  id,
		})
		if err != nil {
			fmt.Printf("Error updating note: %v\n", err)
			return
		}

		fmt.Printf("Note updated successfully: \n\tNote: %v\n\tDescription: %v\n\tCategory: %v\n", *note, description, *categoryName)

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
