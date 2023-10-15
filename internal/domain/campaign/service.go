package campaign

import (
	"github.com/danilocordeirodev/go-email/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaignReq contract.NewCampaignReq) (string, error) {
	campaign, err := NewCampaign(
		newCampaignReq.Name,
		newCampaignReq.Content,
		newCampaignReq.Emails,
	)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)
	if err != nil {
		return "", err
	}
	return campaign.ID, nil
}
