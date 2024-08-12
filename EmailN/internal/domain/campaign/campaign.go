package campaign

import (
	"errors"
	"github.com/rs/xid"
	"time"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contact
}

func NewCampaign(name, content string, emails []string) (*Campaign, error) {

	if name == "" {
		return nil, errors.New("name is required")
	}
	if content == "" {
		return nil, errors.New("content is required")
	}
	if len(emails) == 0 {
		return nil, errors.New("contacts is required")
	}

	contacts := make([]Contact, len(emails))
	for i, v := range emails {
		contacts[i].Email = v
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}, nil
}
