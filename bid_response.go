package gortb


type BidResponse struct {
	ID         string       `json:"id" binding:"required"`
	SeatBids   []SeatBid    `json:"seatbid" binding:"required"`
	BidID      string       `json:"bidid,omitempty"`
	Cur        string       `json:"cur,omitempty" default:"USD"`
	CustomData string       `json:"customdata,omitempty"`
	NBR        int          `json:"nbr,omitempty"`
	Ext        interface{}  `json:"ext,omitempty"`
}