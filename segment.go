package gortb

type Segment struct {
	ID    string      `json:"id,omitempty"`
	Name  string      `json:"name,omitempty"`
	Value string      `json:"value,omitempty"`
	Ext   interface{} `json:"ext,omitempty"`
}
