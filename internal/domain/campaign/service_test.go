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

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}


func (r *repositoryMock) FindAll() []Campaign {
	return nil
}

var (
	newCampaignReq = contract.NewCampaignReq{
		Name:    "Test Y",
		Content: "Body HI",
		Emails:  []string{"test@test.com"},
	}

	service = Service{}
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
