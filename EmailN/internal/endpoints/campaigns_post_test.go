package endpoints

import (
	"bytes"
	"emailn/internal/contract"
	internalmock "emailn/internal/test/internal-mock"
	"encoding/json"
	"fmt"
	assert2 "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CampaignsPost_should_save_new_campaign(t *testing.T) {
	assert := assert2.New(t)
	body := contract.NewCampaign{
		Name:    "Test",
		Content: "Hi everyone",
		Emails:  []string{"teste@teste.com"},
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(newCampaign contract.NewCampaign) bool {
		if newCampaign.Name == body.Name && newCampaign.Content == body.Content {
			return true
		}
		return false
	})).Return("123", nil)
	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest(http.MethodPost, "/", &buf)
	res := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(res, req)

	assert.Equal(http.StatusCreated, status)
	assert.Nil(err)
}

func Test_CampaignsPost_should_inform_error_when_exist(t *testing.T) {
	assert := assert2.New(t)
	body := contract.NewCampaign{
		Name:    "Test",
		Content: "Hi everyone",
		Emails:  []string{"teste@teste.com"},
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest(http.MethodPost, "/", &buf)
	res := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(res, req)

	assert.NotNil(err)
}
