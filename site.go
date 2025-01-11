package gortb

type Site struct {
	ID           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Domain       string   `json:"domain,omitempty"`
	Cat          []string `json:"cat,omitempty"`
	SectionCat   []string `json:"sectioncat,omitempty"`
	PageCat      []string `json:"pagecat,omitempty"`
	Page         string   `json:"page,omitempty"`
	Ref          string   `json:"ref,omitempty"`
	Search       string   `json:"search,omitempty"`
	Mobile       int      `json:"mobile,omitempty"`
	PrivacyPolicy int     `json:"privacypolicy,omitempty"`
	Publisher    *Publisher `json:"publisher,omitempty"`
	Content      *Content   `json:"content,omitempty"`
	Keywords     string    `json:"keywords,omitempty"`
	Ext          interface{} `json:"ext,omitempty"`
}