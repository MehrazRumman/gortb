package gortb

import (
	"errors"
)

// Validation errors for Bid
var (
	ErrMissingBidID = errors.New("bid missing required ID")
	ErrMissingImpID = errors.New("bid missing required impression ID") 
	ErrInvalidPrice = errors.New("bid price must be greater than 0")
	ErrInvalidAPI = errors.New("bid API must be between 1 and 6")
	ErrInvalidProtocol = errors.New("bid protocol must be between 1 and 10")
	ErrInvalidQAGMediaRating = errors.New("bid QAG media rating must be between 1 and 3")
)

type Bid struct {
	ID            string      `json:"id" binding:"required"`
	ImpID         string      `json:"impid" binding:"required"`
	Price         float64     `json:"price" binding:"required"`
	NURL          string      `json:"nurl,omitempty"`
	BURL          string      `json:"burl,omitempty"`
	LURL          string      `json:"lurl,omitempty"`
	AdM           string      `json:"adm,omitempty"`
	AdID          string      `json:"adid,omitempty"`
	ADomain       []string    `json:"adomain,omitempty"`
	Bundle        string      `json:"bundle,omitempty"`
	IURL          string      `json:"iurl,omitempty"`
	CID           string      `json:"cid,omitempty"`
	CRID          string      `json:"crid,omitempty"`
	Tactic        string      `json:"tactic,omitempty"`
	Cat           []string    `json:"cat,omitempty"`
	Attr          []int       `json:"attr,omitempty"`
	API           int         `json:"api,omitempty"`
	Protocol      int         `json:"protocol,omitempty"`
	QAGMediaRating int        `json:"qagmediarating,omitempty"`
	Language      string      `json:"language,omitempty"`
	DealID        string      `json:"dealid,omitempty"`
	W             int         `json:"w,omitempty"`
	H             int         `json:"h,omitempty"`
	WRatio        int         `json:"wratio,omitempty"`
	HRatio        int         `json:"hratio,omitempty"`
	Exp           int         `json:"exp,omitempty"`
	Ext           interface{} `json:"ext,omitempty"`
}

// Validate performs validation on the Bid object according to OpenRTB 2.5 rules
func (b *Bid) Validate() error {
	// Check required ID field
	if b.ID == "" {
		return ErrMissingBidID
	}

	// Check required impression ID field
	if b.ImpID == "" {
		return ErrMissingImpID
	}

	// Check required price field
	if b.Price <= 0 {
		return ErrInvalidPrice
	}

	return nil
}