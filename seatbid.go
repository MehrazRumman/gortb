package gortb

type SeatBid struct {
	Bids  []Bid        `json:"bid" binding:"required"`
	Seat  string       `json:"seat,omitempty"`
	Group int          `json:"group,omitempty" default:"0"`
	Ext   interface{}  `json:"ext,omitempty"`
}