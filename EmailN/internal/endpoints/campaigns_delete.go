package endpoints

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) CampaignDelete(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")

	err := h.CampaignService.Delete(id)
	return nil, http.StatusOK, err
}
