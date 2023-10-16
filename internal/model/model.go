package model

type Product struct {
	Id          int64   `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Rate        float64 `json:"rate"`
}
