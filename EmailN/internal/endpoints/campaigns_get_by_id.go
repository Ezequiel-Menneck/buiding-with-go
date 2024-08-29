package endpoints

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) CampaignGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	campaigns, err := h.CampaignService.GetById(id)
	if err == nil && campaigns == nil {
		return nil, http.StatusNotFound, err
	}
	return campaigns, http.StatusOK, err
}
