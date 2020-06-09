package lead_test

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v5"
	"github.com/splode/go-input-validation-demo/lead"
	"testing"
	"time"
)

func TestJSONToLead(t *testing.T) {
	js := `{"name": "Joe", "email": "joe@example.com", "organization": "Example, Inc.", "message": "I'm interested in learning more about your project.", "phone": "555-555-5555", "newsletter": false, "products": ["hardware"]}`

	var l lead.Lead
	if err := json.Unmarshal([]byte(js), &l); err != nil {
		t.Errorf("failed to unmarshal lead to JSON: %v", err.Error())
	}

	expected := &lead.Lead{
		Name:         "Joe",
		Email:        "joe@example.com",
		Organization: "Example, Inc.",
		Message:      "I'm interested in learning more about your project.",
		Phone:        "555-555-5555",
		Newsletter:   false,
		Products:     []string{"hardware"},
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
	if expected.Phone != l.Phone {
		t.Errorf("expected %v, got %v", expected.Phone, l.Phone)
	}
	if expected.Newsletter != l.Newsletter {
		t.Errorf("expected %v, got %v", expected.Newsletter, l.Newsletter)
	}
	if expected.Products[0] != l.Products[0] {
		t.Errorf("expected %v, got %v", expected.Products[0], l.Products[0])
	}
}

func TestLeadPassesValidation(t *testing.T) {
	js := `{"name": "Joe", "email": "joe@example.com", "organization": "Example, Inc.", "message": "I'm interested in learning more about your project.", "phone": "555-555-5555", "newsletter": false, "products": ["hardware"]}`

	var l lead.Lead
	if err := json.Unmarshal([]byte(js), &l); err != nil {
		t.Errorf("failed to unmarshal lead to JSON: %v", err.Error())
	}
	if err := l.Validate(); err != nil {
		t.Errorf("expected validation on %v, got %v", l, err)
	}
}

func TestLeadNameRequired(t *testing.T) {
	js := `{"email": "joe@example.com", "organization": "Example, Inc.", "message": "I'm interested in learning more about your project.", "phone": "555-555-5555", "newsletter": false}`

	var l lead.Lead
	if err := json.Unmarshal([]byte(js), &l); err != nil {
		t.Errorf("failed to unmarshal lead to JSON: %v", err.Error())
	}
	if err := l.Validate(); err == nil {
		t.Errorf("expected validation error, none received")
	}
}

func TestLeadNameEmoji(t *testing.T) {
	gofakeit.Seed(time.Now().Unix())
	js := fmt.Sprintf(`{"name": "%s", "email": "joe@example.com", "organization": "Example, Inc.", "message": "I'm interested in learning more about your project.", "phone": "555-555-5555", "newsletter": false}`, gofakeit.Emoji())

	var l lead.Lead
	if err := json.Unmarshal([]byte(js), &l); err != nil {
		t.Errorf("failed to unmarshal lead to JSON: %v", err.Error())
	}
	if err := l.Validate(); err == nil {
		t.Errorf("expected validation error, none received")
	}
}

func TestRandomLead(t *testing.T) {
	gofakeit.Seed(time.Now().Unix())
	js := fmt.Sprintf(`{"name": "%s", "email": "%s", "organization": "%s", "message": "%s", "phone": "%s", "newsletter": %t}`, gofakeit.Name(), gofakeit.Email(), gofakeit.Company(), gofakeit.HackerPhrase(), gofakeit.Phone(), gofakeit.Bool())

	var l lead.Lead
	if err := json.Unmarshal([]byte(js), &l); err != nil {
		t.Errorf("failed to unmarshal lead to JSON: %v", err.Error())
	}
	if err := l.Validate(); err != nil {
		t.Errorf("expected validation on %v, got %v", l, err)
	}
}

func TestLeadCharacterLimits(t *testing.T) {
	gofakeit.Seed(time.Now().Unix())
	js := fmt.Sprintf(`{"name": "%s", "email": "%s", "organization": "%s", "message": "%s", "phone": "%s", "newsletter": %t}`, gofakeit.Name(), gofakeit.Email(), gofakeit.Sentence(512), gofakeit.Sentence(512), gofakeit.Phone(), gofakeit.Bool())

	var l lead.Lead
	if err := json.Unmarshal([]byte(js), &l); err != nil {
		t.Errorf("failed to unmarshal lead to JSON: %v", err.Error())
	}
	if err := l.Validate(); err == nil {
		t.Errorf("expected validation error, none received")
	}
}

func TestLeadProducts(t *testing.T) {
	gofakeit.Seed(time.Now().Unix())
	js := fmt.Sprintf(`{"name": "%s", "email": "%s", "organization": "%s", "message": "%s", "phone": "%s", "newsletter": %t, "products": ["hardware", "software"]}`, gofakeit.Name(), gofakeit.Email(), gofakeit.Company(), gofakeit.HackerPhrase(), gofakeit.Phone(), gofakeit.Bool())

	var l lead.Lead
	if err := json.Unmarshal([]byte(js), &l); err != nil {
		t.Errorf("failed to unmarshal lead to JSON: %v", err.Error())
	}
	if err := l.Validate(); err != nil {
		t.Errorf("expected validation on %v, got %v", l, err)
	}
}

func TestLeadProductsInvalid(t *testing.T) {
	gofakeit.Seed(time.Now().Unix())
	js := fmt.Sprintf(`{"name": "%s", "email": "%s", "organization": "%s", "message": "%s", "phone": "%s", "newsletter": %t, "products": ["toast"]}`, gofakeit.Name(), gofakeit.Email(), gofakeit.Company(), gofakeit.HackerPhrase(), gofakeit.Phone(), gofakeit.Bool())

	var l lead.Lead
	if err := json.Unmarshal([]byte(js), &l); err != nil {
		t.Errorf("failed to unmarshal lead to JSON: %v", err.Error())
	}
	if err := l.Validate(); err == nil {
		t.Errorf("expected validation error, none received")
	}
}

func TestPhoneInvalid(t *testing.T) {
	js := `{"name": "", "email": "joe@example.com", "organization": "Example, Inc.", "message": "I'm interested in learning more about your project.", "phone": "5", "newsletter": false}`

	var l lead.Lead
	if err := json.Unmarshal([]byte(js), &l); err != nil {
		t.Errorf("failed to unmarshal lead to JSON: %v", err.Error())
	}
	if err := l.Validate(); err == nil {
		t.Errorf("expected validation error, none received")
	}
}
