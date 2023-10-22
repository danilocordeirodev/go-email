package campaign

import (
	"github.com/danilocordeirodev/go-email/internal/contract"
	internalerrors "github.com/danilocordeirodev/go-email/internal/internalErrors"
)

type Service interface {
	Create(newCampaignReq contract.NewCampaignReq) (string, error)
	GetBy(id string) (*contract.CampaignResponse, error)
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

func (s *ServiceImp) GetBy(id string) (*contract.CampaignResponse, error) {
	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	return &contract.CampaignResponse{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, nil
}
