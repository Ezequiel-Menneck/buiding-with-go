package campaign

import (
	internalerrors "emailn/internal/internal-errors"
	"github.com/rs/xid"
	"time"
)

const (
	Pending  string = "Pending"
	Canceled        = "Canceled"
	Deleted         = "Deleted"
	Started         = "Started"
	Done            = "Done"
)

type Contact struct {
	Id         string `gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:50"`
	CampaignId string `gorm:"size:50"`
}

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024" gorm:"size:1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `gorm:"size:20"`
}

func (c *Campaign) Cancel() {
	c.Status = Canceled
}

func (c *Campaign) Delete() {
	c.Status = Deleted
}

func NewCampaign(name, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	for i, v := range emails {
		contacts[i].Email = v
		contacts[i].Id = xid.New().String()
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
		Status:    Pending,
	}

	err := internalerrors.ValidadeStruct(campaign)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}
