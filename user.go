type User struct {
	ID         string      `json:"id,omitempty"`
	BuyerUID   string      `json:"buyeruid,omitempty"`
	YOB        int         `json:"yob,omitempty"`
	Gender     string      `json:"gender,omitempty"`
	Keywords   string      `json:"keywords,omitempty"`
	CustomData string      `json:"customdata,omitempty"`
	Geo        *Geo        `json:"geo,omitempty"`
	Data       []Data      `json:"data,omitempty"`
	Ext        interface{} `json:"ext,omitempty"`
}