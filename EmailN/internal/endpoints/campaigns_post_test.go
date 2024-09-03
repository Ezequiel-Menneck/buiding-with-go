package endpoints

import (
	"bytes"
	"context"
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

func setup(body contract.NewCampaign, createdByExpected string) (*http.Request, *httptest.ResponseRecorder) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		return nil, nil
	}
	req, _ := http.NewRequest(http.MethodPost, "/", &buf)
	ctx := context.WithValue(req.Context(), "email", createdByExpected)
	req = req.WithContext(ctx)
	res := httptest.NewRecorder()

	return req, res
}

func Test_CampaignsPost_should_save_new_campaign(t *testing.T) {
	assert := assert2.New(t)
	createdByExpected := "teste@teste.com"
	body := contract.NewCampaign{
		Name:    "Test",
		Content: "Hi everyone",
		Emails:  []string{"teste@teste.com"},
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(newCampaign contract.NewCampaign) bool {
		if newCampaign.Name == body.Name && newCampaign.Content == body.Content && newCampaign.CreatedBy == createdByExpected {
			return true
		}
		return false
	})).Return("123", nil)
	handler := Handler{CampaignService: service}

	req, res := setup(body, createdByExpected)

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

	req, res := setup(body, "teste@teste.com")

	_, _, err := handler.CampaignPost(res, req)

	assert.NotNil(err)
}
