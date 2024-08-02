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
)

// getNoteCmd represents the getNote command
var getNoteCmd = &cobra.Command{
	Use:   "getNote",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		note, err := data.GetNoteByName(noteName)
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Não temos nenhuma nota com esse nome")
			return
		}

		category, err := data.FindCategoryNameById(note.CategoryId)
		if err != nil {
			fmt.Println("Algo deu errado ao buscar a categoria", err)
			fmt.Printf("\tNote: %v \n \t Description: %v \n \t", note.NoteName, note.Description)
			return
		}

		fmt.Printf("\tNote: %v \n \t Description: %v \n \t Category: %v \n", note.NoteName, note.Description, category)
		return
	},
}

func init() {
	rootCmd.AddCommand(getNoteCmd)

	getNoteCmd.Flags().StringVarP(&noteName, "note", "n", "", "note to add")
}
