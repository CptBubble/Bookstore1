package models

type ModelBookPrice struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Saleprice   string `json:"price,omitempty"`
}
