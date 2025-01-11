type Deal struct {
	ID         string   `json:"id" binding:"required"`
	BidFloor   float64  `json:"bidfloor,omitempty" default:"0"`
	BidFloorCur string  `json:"bidfloorcur,omitempty" default:"USD"`
	At         int      `json:"at,omitempty"`
	WSeat      []string `json:"wseat,omitempty"`
	WAdmin     []string `json:"wadmin,omitempty"`
	Ext        interface{} `json:"ext,omitempty"`
}