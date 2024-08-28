package endpoints

import (
	"emailn/internal/contract"
	internalmock "emailn/internal/test/mock"
	"errors"
	assert2 "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CampaignsGetById_should_return_campaign(t *testing.T) {
	assert := assert2.New(t)
	campaignResponse := contract.CampaignResponse{
		Id:      "343",
		Name:    "Test",
		Content: "Hi everyone",
		Status:  "Pending",
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("GetById", mock.Anything).Return(&campaignResponse, nil)
	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetById(res, req)

	assert.Equal(http.StatusOK, status)
	assert.Equal(campaignResponse.Id, response.(*contract.CampaignResponse).Id)
	assert.Equal(campaignResponse.Name, response.(*contract.CampaignResponse).Name)
}

func Test_CampaignsGetById_should_return_error_when_something_wrong(t *testing.T) {
	assert := assert2.New(t)
	service := new(internalmock.CampaignServiceMock)
	errExpected := errors.New("something wrong")
	service.On("GetById", mock.Anything).Return(nil, errExpected)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	_, _, errReturned := handler.CampaignGetById(res, req)

	assert.Equal(errExpected.Error(), errReturned.Error())
}
