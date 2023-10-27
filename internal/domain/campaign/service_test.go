package campaign

import (
	"errors"
	"testing"

	"github.com/danilocordeirodev/go-email/internal/contract"
	internalerrors "github.com/danilocordeirodev/go-email/internal/internalErrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Create(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Update(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Get() ([]Campaign, error) {
	return nil, nil
}

func (r *repositoryMock) GetBy(id string) (*Campaign, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaign), nil
}

func (r *repositoryMock) Delete(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaignReq = contract.NewCampaignReq{
		Name:    "Test Y",
		Content: "Body HI",
		Emails:  []string{"test@test.com"},
	}

	service = ServiceImp{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock
	id, err := service.Create(newCampaignReq)

	assert.NotNil(id)
	assert.Nil(err)

}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaignReq{})

	assert.False(errors.Is(internalerrors.ErrInternal, err))

}

func Test_Save_Campaign(t *testing.T) {

	repositoryMock := new(repositoryMock)

	repositoryMock.On("Save", mock.MatchedBy(func(campain *Campaign) bool {
		if campain.Name != newCampaignReq.Name {
			return false
		} else if campain.Content != newCampaignReq.Content {
			return false
		} else if len(campain.Contacts) != len(newCampaignReq.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service.Repository = repositoryMock

	service.Create(newCampaignReq)

	repositoryMock.AssertExpectations(t)
}

func Test_Save_Campaign_ValidateRepositorySave(t *testing.T) {

	assert := assert.New(t)

	repositoryMock := new(repositoryMock)

	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = repositoryMock

	_, err := service.Create(newCampaignReq)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_GetById_ReturnsCampaign(t *testing.T) {

	assert := assert.New(t)

	campaign, _ := NewCampaign(newCampaignReq.Name, newCampaignReq.Content, newCampaignReq.Emails)

	repositoryMock := new(repositoryMock)

	repositoryMock.On("GetBy", mock.Anything).Return(campaign, nil)
	service.Repository = repositoryMock
	campaignReturned, _ := service.GetBy(campaign.ID)

	assert.Equal(campaign.ID, campaignReturned.ID)
}

func Test_GetById_ReturnsError(t *testing.T) {

	assert := assert.New(t)

	campaign, _ := NewCampaign(newCampaignReq.Name, newCampaignReq.Content, newCampaignReq.Emails)

	repositoryMock := new(repositoryMock)

	repositoryMock.On("GetBy", mock.Anything).Return(nil, errors.New("wrong"))
	service.Repository = repositoryMock
	_, err := service.GetBy(campaign.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}
