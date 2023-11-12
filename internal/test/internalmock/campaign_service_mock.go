package internalmock

import (
	"github.com/danilocordeirodev/go-email/internal/contract"
	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (s *CampaignServiceMock) Create(newCampaignReq contract.NewCampaignReq) (string, error) {
	args := s.Called(newCampaignReq)
	return args.String(0), args.Error(1)
}

func (s *CampaignServiceMock) GetBy(id string) (*contract.CampaignResponse, error) {
	args := s.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contract.CampaignResponse), args.Error(1)
}

func (s *CampaignServiceMock) Delete(id string) error {
	args := s.Called(id)
	if args.Error(1) != nil {
		return args.Error(0)
	}
	return args.Error(0)
}
