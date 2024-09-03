package endpoints

import (
	"context"
	internalmock "emailn/internal/test/internal-mock"
	"errors"
	"github.com/go-chi/chi/v5"
	assert2 "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CampaignsStart_200(t *testing.T) {
	assert := assert2.New(t)
	service := new(internalmock.CampaignServiceMock)
	campaignId := "xpto"
	service.On("Start", mock.MatchedBy(func(id string) bool {
		return id == campaignId
	})).Return(nil)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest(http.MethodPatch, "/", nil)
	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add("id", campaignId)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))
	res := httptest.NewRecorder()

	_, status, err := handler.CampaignStart(res, req)

	assert.Equal(http.StatusOK, status)
	assert.Nil(err)
}

func Test_CampaignsStart_Err(t *testing.T) {
	assert := assert2.New(t)
	service := new(internalmock.CampaignServiceMock)
	errExpected := errors.New("something wrong")
	service.On("Start", mock.Anything).Return(errExpected)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest(http.MethodPatch, "/", nil)
	res := httptest.NewRecorder()

	_, _, err := handler.CampaignStart(res, req)

	assert.Equal(errExpected, err)
}
