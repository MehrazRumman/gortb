package gortb

import (
	"errors"
	"fmt"
)

// Validation errors for bid requests
var (
	ErrMissingID = errors.New("bid request missing required ID")
	ErrMissingImp = errors.New("bid request missing required impressions")
	ErrInvalidTest = errors.New("bid request test field must be 0 or 1")
	ErrInvalidAt = errors.New("bid request at (auction type) must be 1 or 2")
	ErrInvalidAllImps = errors.New("bid request allimps field must be 0 or 1")
	ErrSiteAndApp = errors.New("bid request cannot contain both site and app")
)

// BidRequest represents the top-level OpenRTB 2.5 object that contains a bid request.
// The "id" and "imp" fields are required; all others are optional.
// Additional OpenRTB request parameters like site/app, device, user etc. help provide 
// more context for better ad targeting and decision making.
type BidRequest struct {
	// Unique ID of the bid request
	ID       string       `json:"id" binding:"required"`
	// Array of Imp objects representing the impressions offered
	Imp      []Imp        `json:"imp" binding:"required"`
	// Details about the website calling for the impression
	Site     *Site        `json:"site,omitempty"`
	// Details about the application calling for the impression
	App      *App         `json:"app,omitempty"`
	// Device information like type, OS, browser etc
	Device   *Device      `json:"device,omitempty"`
	// Information about the human user of the device
	User     *User        `json:"user,omitempty"`
	// Flag to indicate if this is a test bid request
	Test     int          `json:"test,omitempty" default:"0"`
	// Auction type (1=First Price, 2=Second Price Plus)
	At       int          `json:"at,omitempty" default:"2"`
	// Maximum time in milliseconds for bid response
	TMax     int          `json:"tmax,omitempty"`
	// Whitelist of buyer seats allowed to bid
	WSeat    []string     `json:"wseat,omitempty"`
	// Block list of buyer seats not allowed to bid
	BSeat    []string     `json:"bseat,omitempty"`
	// Flag to indicate all impressions are available
	AllImps  int          `json:"allimps,omitempty" default:"0"`
	// Array of allowed currencies for bids
	Cur      []string     `json:"cur,omitempty"`
	// Whitelist of languages for creatives
	WLang    []string     `json:"wlang,omitempty"`
	// Blocked advertiser categories using IAB taxonomy
	BCat     []string     `json:"bcat,omitempty"`
	// Block list of advertiser domains
	BAdv     []string     `json:"badv,omitempty"`
	// Block list of applications
	BApp     []string     `json:"bapp,omitempty"`
	// Object describing the source of the bid request
	Source   *Source      `json:"source,omitempty"`
	// Object containing any legal, governmental, or industry regulations
	Regs     *Regs        `json:"regs,omitempty"`
	// Placeholder for exchange-specific extensions
	Ext      interface{}  `json:"ext,omitempty"`
}

// Validate performs validation on the BidRequest object according to OpenRTB 2.5 rules
func (br *BidRequest) Validate() error {
	// Check required fields
	if br.ID == "" {
		return ErrMissingID
	}
	if len(br.Imp) == 0 {
		return ErrMissingImp
	}

	// Cannot have both site and app
	if br.Site != nil && br.App != nil {
		return ErrSiteAndApp
	}

	// Must have either site or app
	if br.Site == nil && br.App == nil {
		return errors.New("bid request must contain either site or app")
	}

	// Validate test field
	if br.Test != 0 && br.Test != 1 {
		return ErrInvalidTest
	}

	// Validate auction type
	if br.At != 0 && br.At != 1 && br.At != 2 {
		return ErrInvalidAt
	}

	// Validate allimps field
	if br.AllImps != 0 && br.AllImps != 1 {
		return ErrInvalidAllImps
	}

	// Validate each impression
	for _, imp := range br.Imp {
		if err := imp.Validate(); err != nil {
			return fmt.Errorf("invalid impression: %v", err)
		}
	}

	return nil
}












