package campaign

import (
	internalerrors "emailn/internal/internal-errors"
	"github.com/rs/xid"
	"time"
)

const (
	Pending string = "Pending"
	Started        = "Started"
	Done           = "Done"
)

type Contact struct {
	Id         string
	Email      string `validate:"email"`
	CampaignId string
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string
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
