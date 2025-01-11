type App struct {
	ID            string   `json:"id" binding:"required"`
	Name          string   `json:"name,omitempty"`
	Bundle        string   `json:"bundle,omitempty"`
	Domain        string   `json:"domain,omitempty"`
	StoreURL      string   `json:"storeurl,omitempty"`
	Cat           []string `json:"cat,omitempty"`
	SectionCat    []string `json:"sectioncat,omitempty"`
	PageCat       []string `json:"pagecat,omitempty"`
	Ver           string   `json:"ver,omitempty"`
	PrivacyPolicy int      `json:"privacypolicy,omitempty"`
	Paid          int      `json:"paid,omitempty"`
	Publisher     *Publisher `json:"publisher,omitempty"`
	Content       *Content   `json:"content,omitempty"`
	Keywords      string    `json:"keywords,omitempty"`
	Ext           interface{} `json:"ext,omitempty"`
}