package gortb

import (
	"errors"
)

// Validation errors for Deal
var (
	ErrMissingDealID = errors.New("deal missing required ID")
)

type Deal struct {
	ID         string   `json:"id" binding:"required"`
	BidFloor   float64  `json:"bidfloor,omitempty" default:"0"`
	BidFloorCur string  `json:"bidfloorcur,omitempty" default:"USD"`
	At         int      `json:"at,omitempty"`
	WSeat      []string `json:"wseat,omitempty"`
	WAdmin     []string `json:"wadmin,omitempty"`
	Ext        interface{} `json:"ext,omitempty"`
}

// Validate performs validation on the Deal object according to OpenRTB 2.5 rules
func (d *Deal) Validate() error {
	// Check required ID field
	if d.ID == "" {
		return ErrMissingDealID
	}
	return nil
}