package main

// Lead represents the lead form POST data.
type Lead struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Organization string `json:"organization"`
	Message      string `json:"message"`
	Referrer     string `json:"referrer"`
}
