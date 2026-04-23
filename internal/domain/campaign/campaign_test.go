package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "teste1"
	content := "teste2"
	emails := []string{"filipe@e.com", "gustavo@e.com", "alex@e.com"}

	campaign := NewCampaign(name, content, emails)
	if campaign.ID != "1" {
		t.Errorf("Expected 1")
	} else if campaign.Content != content {
		t.Errorf("Expected correct content")
	} else if campaign.Name != name {
		t.Errorf("Expected correct content")
	} else if len(campaign.Contacts) != len(emails) {
		t.Errorf("Expected correct emails")
	}
}
