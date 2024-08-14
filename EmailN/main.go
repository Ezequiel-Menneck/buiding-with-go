package main

import (
	"emailn/internal/domain/campaign"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func main() {
	c := campaign.Campaign{}
	validate := validator.New()
	err := validate.Struct(c)
	if err == nil {
		fmt.Println("Nenhum erro")
	} else {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		for _, validationError := range validationErrors {
			fmt.Println(validationError.Error())
		}
	}
}
