package models

type Book struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Notes 	    string `json:"notes"`
	ImageUrl	string `json:"imageUrl"`
	BuyURL      string `json:"buyUrl"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}
