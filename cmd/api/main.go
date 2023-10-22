package main

import (
	"net/http"

	"github.com/danilocordeirodev/go-email/internal/domain/campaign"
	"github.com/danilocordeirodev/go-email/internal/endpoints"
	"github.com/danilocordeirodev/go-email/internal/infrastructure/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{},
	}

	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaigns/{id}", endpoints.HandlerError(handler.CampaignGetById))
	http.ListenAndServe(":3000", r)
}
