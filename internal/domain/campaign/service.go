package campaign

import (
	"github.com/danilocordeirodev/go-email/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaignReq contract.NewCampaignReq) (string, error) {
	campaign, _ := NewCampaign(
		newCampaignReq.Name,
		newCampaignReq.Content,
		newCampaignReq.Emails,
	)
	s.Repository.Save(campaign)
	return campaign.ID, nil
}
