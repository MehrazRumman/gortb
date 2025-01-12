package gortb

type PMP struct {
	PrivateAuction int         `json:"private_auction,omitempty" default:"0"`
	Deals          []Deal      `json:"deals,omitempty"`
	Ext            interface{} `json:"ext,omitempty"`
}
