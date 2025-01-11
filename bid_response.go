package gortb

import (
	"errors"
)

// Validation errors for BidResponse
var (
	ErrMissingResponseID = errors.New("bid response missing required ID")
	ErrMissingSeatBids = errors.New("bid response missing required seat bids")
	ErrInvalidNBR = errors.New("bid response NBR must be between 0 and 10")
)

type BidResponse struct {
	ID         string       `json:"id" binding:"required"`
	SeatBids   []SeatBid    `json:"seatbid" binding:"required"`
	BidID      string       `json:"bidid,omitempty"`
	Cur        string       `json:"cur,omitempty" default:"USD"`
	CustomData string       `json:"customdata,omitempty"`
	NBR        int          `json:"nbr,omitempty"`
	Ext        interface{}  `json:"ext,omitempty"`
}

// Validate performs validation on the BidResponse object according to OpenRTB 2.5 rules
func (br *BidResponse) Validate() error {
	// Check required ID field
	if br.ID == "" {
		return ErrMissingResponseID
	}

	// Check required SeatBids field
	if len(br.SeatBids) == 0 {
		return ErrMissingSeatBids
	}

	// Validate NBR if present (values 0-10 are valid)
	if br.NBR != 0 && (br.NBR < 0 || br.NBR > 10) {
		return ErrInvalidNBR
	}

	// Validate each SeatBid
	for _, sb := range br.SeatBids {
		if err := sb.Validate(); err != nil {
			return err
		}
	}

	return nil
}