package database

import (
	"emailn/internal/domain/campaign"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	tx := c.Db.Create(campaign)
	return tx.Error
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	tx := c.Db.Find(&campaigns)

	return campaigns, tx.Error
}

func (c *CampaignRepository) GetById(id string) (*campaign.Campaign, error) {
	var campaignToReturn campaign.Campaign
	tx := c.Db.First(&campaignToReturn, "id = ?", id)
	return &campaignToReturn, tx.Error
}
