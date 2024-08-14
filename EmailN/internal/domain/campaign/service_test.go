package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	assert2 "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test CarroDeSom",
		Content: "Body eeeeee",
		Emails:  []string{"test@test.com"},
	}
	service = Service{}
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert2.New(t)
	repositoryMock := new(RepositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
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
	repositoryMock := new(RepositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
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
	repositoryMock := new(RepositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}
