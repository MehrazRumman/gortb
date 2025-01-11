type Producer struct {
	ID    string   `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Cat   []string `json:"cat,omitempty"`
	Domain string `json:"domain,omitempty"`
	Ext   interface{} `json:"ext,omitempty"`
}