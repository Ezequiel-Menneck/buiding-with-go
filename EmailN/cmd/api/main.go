package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"emailn/internal/infrastructure/mail"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.NewDb()
	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
		SendMail:   mail.SendMail,
	}
	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)
		r.Post("/", endpoints.HandlerError(handler.CampaignPost))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetById))
		r.Delete("/delete/{id}", endpoints.HandlerError(handler.CampaignDelete))
		r.Patch("/start/{id}", endpoints.HandlerError(handler.CampaignStart))
	})

	err = http.ListenAndServe(":8081", r)
	if err != nil {
		return
	}
}
