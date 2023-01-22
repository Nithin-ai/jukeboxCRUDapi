package models

type Song struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Artist      string `json:"artist"`
	Description string `json:"description"`
}
