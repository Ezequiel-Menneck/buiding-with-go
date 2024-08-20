package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	handler := endpoints.Handler{}
	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{},
	}
	handler.CampaignService = &campaignService
	r.Post("/campaign", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaign", endpoints.HandlerError(handler.CampaignGet))

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
