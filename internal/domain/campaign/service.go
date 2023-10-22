package campaign

import (
	"github.com/danilocordeirodev/go-email/internal/contract"
	internalerrors "github.com/danilocordeirodev/go-email/internal/internalErrors"
)

type Service interface {
	Create(newCampaignReq contract.NewCampaignReq) (string, error)
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaignReq contract.NewCampaignReq) (string, error) {
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
		return "", internalerrors.ErrInternal
	}
	return campaign.ID, nil
}
