package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Bodyuu"
	contacts = []string{"email@e.com", "email2@e.com"}
	fake = faker.New()
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
	assert.NotNil(campaign.CreatedOn)

}

func Test_NewCampaign_CreatedOnShouldBeNow(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)
	now := time.Now().Add(-time.Minute)

	assert.Greater(campaign.CreatedOn, now)

}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("Name is required with min 5", err.Error())

}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)
	
	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	assert.Equal("Name is required with max 24", err.Error())

}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("Content is required with min 5", err.Error())

}
func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts)

	assert.Equal("Content is required with max 1024", err.Error())

}
func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("Contacts is required with min 1", err.Error())

}

func Test_NewCampaign_MustValidateContactsWithValidEmail(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalid"})

	assert.Equal("Email is invalid.", err.Error())

}

func Test_NewCampaign_MustStatusStartWithPending(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(Pending, campaign.Status)

}




