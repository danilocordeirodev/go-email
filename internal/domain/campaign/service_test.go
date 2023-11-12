package campaign_test

import (
	"errors"
	"testing"

	"github.com/danilocordeirodev/go-email/internal/contract"
	"github.com/danilocordeirodev/go-email/internal/domain/campaign"
	internalerrors "github.com/danilocordeirodev/go-email/internal/internalErrors"
	internalmock "github.com/danilocordeirodev/go-email/internal/test/internalmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	newCampaignReq = contract.NewCampaignReq{
		Name:    "Test Y",
		Content: "Body HI",
		Emails:  []string{"test@test.com"},
	}

	service = campaign.ServiceImp{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(nil)
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

	repositoryMock := new(internalmock.CampaignRepositoryMock)

	repositoryMock.On("Create", mock.MatchedBy(func(campain *campaign.Campaign) bool {
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

	repositoryMock := new(internalmock.CampaignRepositoryMock)

	repositoryMock.On("Create", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = repositoryMock

	_, err := service.Create(newCampaignReq)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_GetById_ReturnsCampaign(t *testing.T) {

	assert := assert.New(t)

	campaign, _ := campaign.NewCampaign(newCampaignReq.Name, newCampaignReq.Content, newCampaignReq.Emails)

	repositoryMock := new(internalmock.CampaignRepositoryMock)

	repositoryMock.On("GetBy", mock.Anything).Return(campaign, nil)
	service.Repository = repositoryMock
	campaignReturned, _ := service.GetBy(campaign.ID)

	assert.Equal(campaign.ID, campaignReturned.ID)
}

func Test_GetById_ReturnsError(t *testing.T) {

	assert := assert.New(t)

	campaign, _ := campaign.NewCampaign(newCampaignReq.Name, newCampaignReq.Content, newCampaignReq.Emails)

	repositoryMock := new(internalmock.CampaignRepositoryMock)

	repositoryMock.On("GetBy", mock.Anything).Return(nil, errors.New("wrong"))
	service.Repository = repositoryMock
	_, err := service.GetBy(campaign.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}

func Test_Delete_ReturnsRecordNotFound_when_campaign_does_not_exist(t *testing.T) {

	assert := assert.New(t)
	campaignIdInvalid := "invalid"
	repositoryMock := new(internalmock.CampaignRepositoryMock)

	repositoryMock.On("GetBy", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
	service.Repository = repositoryMock
	err := service.Delete(campaignIdInvalid)

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())
}

func Test_Delete_ReturnsInvalidStatus_when_campaign_status_not_equals_pending(t *testing.T) {

	assert := assert.New(t)
	campaign := &campaign.Campaign{ID: "1", Status: campaign.Started}
	repositoryMock := new(internalmock.CampaignRepositoryMock)

	repositoryMock.On("GetBy", mock.Anything).Return(campaign, nil)
	service.Repository = repositoryMock
	err := service.Delete(campaign.ID)

	assert.Equal("Campaign status invalid", err.Error())
}


func Test_Delete_ReturnsInternalErro_when_delete_has_problem(t *testing.T) {

	assert := assert.New(t)
	campaignFound, _ := campaign.NewCampaign("Test 1", "body TT", []string{"test@twww.com.br"})
	repositoryMock := new(internalmock.CampaignRepositoryMock)

	repositoryMock.On("GetBy", mock.Anything).Return(campaignFound, nil)
	repositoryMock.On("Delete", mock.MatchedBy(
		func (campain *campaign.Campaign) bool {
			return campaignFound == campain
		})).Return(errors.New("error to delete campaign"))
	service.Repository = repositoryMock
	err := service.Delete(campaignFound.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}

func Test_Delete_ReturnsNil_when_delete_has_success(t *testing.T) {

	assert := assert.New(t)
	campaignFound, _ := campaign.NewCampaign("Test 1", "body TT", []string{"test@twww.com.br"})
	repositoryMock := new(internalmock.CampaignRepositoryMock)

	repositoryMock.On("GetBy", mock.Anything).Return(campaignFound, nil)
	repositoryMock.On("Delete", mock.MatchedBy(
		func (campain *campaign.Campaign) bool {
			return campaignFound == campain
		})).Return(nil)
	service.Repository = repositoryMock
	err := service.Delete(campaignFound.ID)

	assert.Nil(err)
}