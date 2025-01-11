type Data struct {
	ID       string      `json:"id,omitempty"`
	Name     string      `json:"name,omitempty"`
	Segment  []Segment   `json:"segment,omitempty"`
	Ext      interface{} `json:"ext,omitempty"`
}