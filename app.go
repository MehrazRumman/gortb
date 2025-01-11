package gortb

import (
	"errors"
)

// Validation errors for App
var (
	ErrMissingAppID = errors.New("app missing required ID")
	ErrInvalidPrivacyPolicy = errors.New("app privacy policy must be 0 or 1")
	ErrInvalidPaid = errors.New("app paid field must be 0 or 1")
)

type App struct {
	ID            string   `json:"id" binding:"required"`
	Name          string   `json:"name,omitempty"`
	Bundle        string   `json:"bundle,omitempty"`
	Domain        string   `json:"domain,omitempty"`
	StoreURL      string   `json:"storeurl,omitempty"`
	Cat           []string `json:"cat,omitempty"`
	SectionCat    []string `json:"sectioncat,omitempty"`
	PageCat       []string `json:"pagecat,omitempty"`
	Ver           string   `json:"ver,omitempty"`
	PrivacyPolicy int      `json:"privacypolicy,omitempty"`
	Paid          int      `json:"paid,omitempty"`
	Publisher     *Publisher `json:"publisher,omitempty"`
	Content       *Content   `json:"content,omitempty"`
	Keywords      string    `json:"keywords,omitempty"`
	Ext           interface{} `json:"ext,omitempty"`
}

func (a *App) Validate() error {
	// Check required fields
	if a.ID == "" {
		return ErrMissingAppID
	}

	// Validate privacy policy field
	if a.PrivacyPolicy != 0 && a.PrivacyPolicy != 1 {
		return ErrInvalidPrivacyPolicy
	}

	// Validate paid field
	if a.Paid != 0 && a.Paid != 1 {
		return ErrInvalidPaid
	}
	
	return nil
}
