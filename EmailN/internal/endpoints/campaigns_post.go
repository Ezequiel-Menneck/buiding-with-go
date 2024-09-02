package endpoints

import (
	"emailn/internal/contract"
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaign
	err := render.DecodeJSON(r.Body, &request)
	email := r.Context().Value("email").(string)
	request.CreatedBy = email
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	id, err := h.CampaignService.Create(request)
	return map[string]string{"id": id}, http.StatusCreated, err
}
