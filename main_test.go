package main_test

import (
	"encoding/json"
	"github.com/splode/go-input-validation-demo"
	"testing"
)

func TestJSONtoLead(t *testing.T) {
	js := `{"name": "Joe", "email": "joe@example.com", "organization": "Example, Inc.", "message": "I'm interested in learning more about your project.", "referrer": "search engine"}`

	var lead main.Lead
	if err := json.Unmarshal([]byte(js), &lead); err != nil {
		t.Errorf("failed to unmarshal lead to JSON: %v", err.Error())
	}

	expected := &main.Lead{
		Name:         "Joe",
		Email:        "joe@example.com",
		Organization: "Example, Inc.",
		Message:      "I'm interested in learning more about your project.",
		Referrer:     "search engine",
	}

	if expected.Name != lead.Name {
		t.Errorf("expected %v, got %v", expected.Name, lead.Name)
	}
	if expected.Email != lead.Email {
		t.Errorf("expected %v, got %v", expected.Email, lead.Email)
	}
	if expected.Organization != lead.Organization {
		t.Errorf("expected %v, got %v", expected.Organization, lead.Organization)
	}
	if expected.Message != lead.Message {
		t.Errorf("expected %v, got %v", expected.Message, lead.Message)
	}
	if expected.Referrer != lead.Referrer {
		t.Errorf("expected %v, got %v", expected.Referrer, lead.Referrer)
	}
}
