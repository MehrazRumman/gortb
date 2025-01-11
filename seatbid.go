package gortb

import(
	"errors"
)

var (
	ErrInvalidBids = errors.New("seatbid must contain at least one bid")
)

type SeatBid struct {
	Bids  []Bid        `json:"bid" binding:"required"`
	Seat  string       `json:"seat,omitempty"`
	Group int          `json:"group,omitempty" default:"0"`
	Ext   interface{}  `json:"ext,omitempty"`
}


func (sb *SeatBid) Validate() error {
	if len(sb.Bids) == 0 {
		return ErrInvalidBids
	}

	for _, bid := range sb.Bids {
		if err := bid.Validate(); err != nil {
			return err
		}
	}

	return nil
}
