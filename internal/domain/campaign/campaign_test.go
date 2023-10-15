package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
	assert.NotNil(campaign.CreatedOn)

}

func Test_NewCampaign_CreatedOnIsNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)
	now := time.Now().Add(-time.Minute)
	
	assert.Greater(campaign.CreatedOn, now)


}
