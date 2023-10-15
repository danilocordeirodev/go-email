package campaign

import (
	"testing"

	"github.com/danilocordeirodev/go-email/internal/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign )
	return args.Error(0)
}


func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	newCampaignReq := contract.NewCampaignReq{
		Name: "Test Y",
		Content: "Body",
		Emails: []string{"test@test.com"},
	}

	id, err := service.Create(newCampaignReq)

	assert.NotNil(id)
	assert.Nil(err)
	
}

func Test_Save_Campaign(t *testing.T) {
	
	repositoryMock := new(repositoryMock)
	newCampaignReq := contract.NewCampaignReq{
		Name: "Test Y",
		Content: "Body",
		Emails: []string{"test@test.com"},
	}
	repositoryMock.On("Save", mock.MatchedBy(func (campain *Campaign) bool {
		if campain.Name != newCampaignReq.Name {
			return false
		} else if campain.Content != newCampaignReq.Content {
			return false
		} else if len(campain.Contacts) != len(newCampaignReq.Emails) {
			return false
		}
		
		return true
	})).Return(nil)
	service := Service{Repository: repositoryMock}
	

	service.Create(newCampaignReq)

	repositoryMock.AssertExpectations(t)
}