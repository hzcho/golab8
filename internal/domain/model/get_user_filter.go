package model

type GetUserFilter struct {
	Name  string `json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
	Page  int    `json:"page,omitempty"`
	Limit int    `json:"limit,omitempty"`
}
