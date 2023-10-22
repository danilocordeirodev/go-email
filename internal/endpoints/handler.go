package endpoints

import "github.com/danilocordeirodev/go-email/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}