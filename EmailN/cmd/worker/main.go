package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"
	"emailn/internal/infrastructure/mail"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.NewDb()
	repository := &database.CampaignRepository{Db: db}
	campaignService := campaign.ServiceImp{
		Repository: repository,
		SendMail:   mail.SendMail,
	}

	for {
		campaigns, err := repository.GetCampaignsToBeSent()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(campaigns)
		for _, campaignToUpdate := range campaigns {
			campaignService.SendEmailAndUpdateStatus(&campaignToUpdate)
		}

		time.Sleep(time.Second * 10)
	}

}
