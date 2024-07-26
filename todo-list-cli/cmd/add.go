/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"todo-list-cli/data"
)

//type AddActivity struct {
//	ActivityName string
//	Category string
//	StartsAt time.Time
//	EndsAt time.Time
//}

var noteName string
var description string

//var category string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := data.GetNoteByName(noteName)
		if err != nil {
			fmt.Println("something went wrong trying to get note by name")
			return
		}

		//if name {
		//	fmt.Println("note already exists")
		//	return
		//}

		fmt.Println(name)

		//id, err := data.FindCategoryByName(category)
		//if err != nil {
		//	fmt.Errorf("something went wrong trying to get category")
		//	return
		//}
		//
		//if id == 0 {
		//	id, err = data.CreateCategory(category)
		//	if err != nil {
		//		fmt.Errorf("something went wrong trying to create category")
		//		return
		//	}
		//}

		note, err := data.InsertNote(data.Note{
			NoteName:    "carrr",
			Description: "cwweedd",
			CategoryId:  1,
		})
		if err != nil {
			return
		}

		fmt.Printf("note added successfully: %v", note)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&noteName, "note", "n", "", "note to add")
	addCmd.Flags().StringVarP(&description, "description", "d", "", "description of your note")
	addCmd.Flags().StringVarP(&category, "category", "c", "", "category of your note")
}
