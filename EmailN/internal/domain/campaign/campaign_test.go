package campaign

import (
	"github.com/jaswdr/faker"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	name      = "Campaign X"
	content   = "Body eeee"
	contacts  = []string{"email1@example.com", "email2@example.com"}
	fake      = faker.New()
	createdBy = "teste@teste.com"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert2.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
	assert.Equal(campaign.CreatedBy, createdBy)
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert2.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_MustStatusStartWithPending(t *testing.T) {
	assert := assert2.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.Equal(Pending, campaign.Status)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert2.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assert := assert2.New(t)

	_, err := NewCampaign("", content, contacts, createdBy)

	assert.Equal("name is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	assert := assert2.New(t)
	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts, createdBy)

	assert.Equal("name is required with max 24", err.Error())
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert2.New(t)

	_, err := NewCampaign(name, "", contacts, createdBy)

	assert.Equal("content is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert2.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1200), contacts, createdBy)

	assert.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	assert := assert2.New(t)

	_, err := NewCampaign(name, content, []string{}, createdBy)

	assert.Equal("contacts is required with min 1", err.Error())
}

func Test_NewCampaign_MustValidateContactsInvalidEmail(t *testing.T) {
	assert := assert2.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalid"}, createdBy)

	assert.Equal("email is invalid", err.Error())
}

func Test_NewCampaign_MustValidateCreatedBy(t *testing.T) {
	assert := assert2.New(t)

	_, err := NewCampaign(name, content, contacts, "")

	assert.Equal("createdby is invalid", err.Error())
}
