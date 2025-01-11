package gortb

type Source struct {
	FD     int         `json:"fd,omitempty"`
	TID    string      `json:"tid,omitempty"`
	PChain string      `json:"pchain,omitempty"`
	Ext    interface{} `json:"ext,omitempty"`
}
