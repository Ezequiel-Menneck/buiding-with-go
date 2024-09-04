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
	Fail            = "Fail"
)

type Contact struct {
	Id         string `gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:50"`
	CampaignId string `gorm:"size:50"`
}

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50;not null"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100;not null"`
	CreatedOn time.Time `validate:"required" gorm:"not null"`
	UpdatedOn time.Time
	Content   string    `validate:"min=5,max=1024" gorm:"size:1024;not null"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `gorm:"size:20;not null"`
	CreatedBy string    `validate:"email" gorm:"size:50;not null"`
}

func (c *Campaign) Fail() {
	c.Status = Fail
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Done() {
	c.Status = Done
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Cancel() {
	c.Status = Canceled
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Delete() {
	c.Status = Deleted
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Started() {
	c.Status = Started
	c.UpdatedOn = time.Now()
}

func NewCampaign(name, content string, emails []string, createdBy string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	for i, v := range emails {
		contacts[i].Email = v
		contacts[i].Id = xid.New().String()
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
		Status:    Pending,
		CreatedBy: createdBy,
	}

	err := internalerrors.ValidadeStruct(campaign)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}
