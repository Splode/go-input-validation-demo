package lead_test

import (
	"encoding/json"
	"github.com/splode/go-input-validation-demo/lead"
	"testing"
)

func TestJSONtoLead(t *testing.T) {
	js := `{"name": "Joe", "email": "joe@example.com", "organization": "Example, Inc.", "message": "I'm interested in learning more about your project.", "referrer": "search engine"}`

	var l lead.Lead
	if err := json.Unmarshal([]byte(js), &l); err != nil {
		t.Errorf("failed to unmarshal lead to JSON: %v", err.Error())
	}

	expected := &lead.Lead{
		Name:         "Joe",
		Email:        "joe@example.com",
		Organization: "Example, Inc.",
		Message:      "I'm interested in learning more about your project.",
		Referrer:     "search engine",
	}

	if expected.Name != l.Name {
		t.Errorf("expected %v, got %v", expected.Name, l.Name)
	}
	if expected.Email != l.Email {
		t.Errorf("expected %v, got %v", expected.Email, l.Email)
	}
	if expected.Organization != l.Organization {
		t.Errorf("expected %v, got %v", expected.Organization, l.Organization)
	}
	if expected.Message != l.Message {
		t.Errorf("expected %v, got %v", expected.Message, l.Message)
	}
	if expected.Referrer != l.Referrer {
		t.Errorf("expected %v, got %v", expected.Referrer, l.Referrer)
	}
}

func TestLeadNameRequired(t *testing.T) {
	js := `{"name": "", "email": "joe@example.com", "organization": "Example, Inc.", "message": "I'm interested in learning more about your project.", "referrer": "search engine"}`

	var l lead.Lead
	if err := json.Unmarshal([]byte(js), &l); err != nil {
		t.Errorf("failed to unmarshal lead to JSON: %v", err.Error())
	}
	if err := l.Validate(); err == nil {
		t.Errorf("expected validation error, none received")
	}
}
