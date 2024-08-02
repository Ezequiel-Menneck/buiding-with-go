package promptui_utils

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"todo-list-cli/data"
)

func PrintPromptUiCategories() (*string, error) {
	categories, err := data.FindAllCategories()
	if err != nil {
		return nil, err
	}

	prompt := promptui.SelectWithAdd{
		Label:    "Select The Category",
		Items:    categories,
		AddLabel: "Add another Category",
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	return &result, nil
}

func PrintPrompUiNotes() (*string, error) {
	notes, err := data.FindAllNotesName()
	if err != nil {
		return nil, err
	}

	prompt := promptui.Select{
		Label: "Select The Note",
		Items: notes,
	}

	_, reuslt, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	return &reuslt, nil
}

func UserInputToUpdateNote(message string) string {
	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Note %v", message),
		Validate:  validate,
		Templates: templates,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}
