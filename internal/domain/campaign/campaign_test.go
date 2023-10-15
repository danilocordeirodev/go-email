package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)

	if campaign.ID != "1" {
		t.Errorf("expected 1")
	} else if campaign.Name != name {
		t.Errorf("expected correct name")
	} else if campaign.Content != content {
		t.Errorf("expected correct content")
	} else if len(campaign.Contacts) != len(contacts) {
		t.Errorf("expected correct contacts")
	}
}
