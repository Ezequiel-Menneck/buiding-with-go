package campaign_test

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	internalerrors "emailn/internal/internal-errors"
	"emailn/internal/test/internal-mock"
	"errors"
	assert2 "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
)

var (
	newCampaign = contract.NewCampaign{
		Name:      "Test CarroDeSom",
		Content:   "Body eeeeee",
		Emails:    []string{"test@test.com"},
		CreatedBy: "teste@teste.com",
	}
	service = campaign.ServiceImp{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert2.New(t)
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert2.New(t)
	_, err := service.Create(contract.NewCampaign{})

	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)
	service.Repository = repositoryMock

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert2.New(t)
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_GetById_ReturnCampaign(t *testing.T) {
	assert := assert2.New(t)
	campaignMock, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == campaignMock.ID
	})).Return(campaignMock, nil)

	campaignReturned, _ := service.GetById(campaignMock.ID)

	assert.Equal(campaignMock.ID, campaignReturned.Id)
	assert.Equal(campaignMock.Name, campaignReturned.Name)
	assert.Equal(campaignMock.Content, campaignReturned.Content)
	assert.Equal(campaignMock.Status, campaignReturned.Status)
	assert.Equal(campaignMock.CreatedBy, campaignReturned.CreatedBy)
}

func Test_GetById_ReturnErrorWhenSomethingWrongExist(t *testing.T) {
	assert := assert2.New(t)
	campaignMock, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("GetById", mock.Anything).Return(nil, errors.New("something wrong"))

	_, err := service.GetById(campaignMock.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}

func Test_Delete_ReturnRecordNotFound_when_campaign_does_not_exist(t *testing.T) {
	assert := assert2.New(t)
	campaignIdInvalid := "invalid"
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("GetById", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	err := service.Delete(campaignIdInvalid)

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())
}

func Test_Delete_ReturnStatusInvalid_when_campaign_has_status_not_equals_pending(t *testing.T) {
	assert := assert2.New(t)
	campaignMock := &campaign.Campaign{
		ID:     "1",
		Name:   "Carrodesom",
		Status: campaign.Started,
	}
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("GetById", mock.Anything).Return(campaignMock, nil)

	err := service.Delete(campaignMock.ID)

	assert.Equal("campaign status invalid", err.Error())
}

func Test_Delete_ReturnInternalError_when_delete_has_problem(t *testing.T) {
	assert := assert2.New(t)
	campaignMock, _ := campaign.NewCampaign("Test 1", "Some Content", []string{"email@email.com"}, newCampaign.CreatedBy)
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("GetById", mock.Anything).Return(campaignMock, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaign == campaignMock
	})).Return(errors.New("error to delete campaign"))

	err := service.Delete(campaignMock.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}

func Test_Delete_ReturnNil_when_delete_has_success(t *testing.T) {
	assert := assert2.New(t)
	campaignMock, _ := campaign.NewCampaign("Test 1", "Some Content", []string{"email@email.com"}, newCampaign.CreatedBy)
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("GetById", mock.Anything).Return(campaignMock, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaign == campaignMock
	})).Return(nil)

	err := service.Delete(campaignMock.ID)

	assert.Nil(err)
}

func Test_Start_ReturnRecordNotFound_when_campaign_does_not_exist(t *testing.T) {
	assert := assert2.New(t)
	campaignIdInvalid := "invalid"
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("GetById", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	err := service.Start(campaignIdInvalid)

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())
}

func Test_Start_ReturnStatusInvalid_when_campaign_has_status_not_equals_pending(t *testing.T) {
	assert := assert2.New(t)
	campaignMock := &campaign.Campaign{
		ID:     "1",
		Name:   "Carrodesom",
		Status: campaign.Started,
	}
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("GetById", mock.Anything).Return(campaignMock, nil)

	err := service.Start(campaignMock.ID)

	assert.Equal("campaign status invalid", err.Error())
}

func Test_Start_should_send_mail(t *testing.T) {
	assert := assert2.New(t)
	campaignMock := &campaign.Campaign{
		ID:     "1",
		Name:   "Carrodesom",
		Status: campaign.Pending,
	}
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("GetById", mock.Anything).Return(campaignMock, nil)
	emailWasSent := false
	sendMail := func(campaign *campaign.Campaign) error {
		if campaign.ID == campaignMock.ID {
			emailWasSent = true
		}
		return nil
	}
	service.SendMail = sendMail

	_ = service.Start(campaignMock.ID)

	assert.True(emailWasSent)
}

func Test_Start_ReturnError_when_func_SendMail_fail(t *testing.T) {
	assert := assert2.New(t)
	campaignMock := &campaign.Campaign{
		ID:     "1",
		Name:   "Carrodesom",
		Status: campaign.Pending,
	}
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("GetById", mock.Anything).Return(campaignMock, nil)
	sendMail := func(campaign *campaign.Campaign) error {
		return errors.New("error to send mail")
	}
	service.SendMail = sendMail

	err := service.Start(campaignMock.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}

func Test_Start_ReturnNil_when_updated_to_done(t *testing.T) {
	assert := assert2.New(t)
	campaignMock := &campaign.Campaign{
		ID:     "1",
		Name:   "Carrodesom",
		Status: campaign.Pending,
	}
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("GetById", mock.Anything).Return(campaignMock, nil)
	repositoryMock.On("Update", mock.MatchedBy(func(campaignToUpdate *campaign.Campaign) bool {
		return campaignToUpdate.ID == campaignMock.ID && campaignToUpdate.Status == campaign.Done
	})).Return(nil)
	sendMail := func(campaign *campaign.Campaign) error {
		return nil
	}
	service.SendMail = sendMail

	_ = service.Start(campaignMock.ID)

	assert.Equal(campaignMock.Status, campaign.Done)

}
