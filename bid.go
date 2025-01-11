package gortb


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