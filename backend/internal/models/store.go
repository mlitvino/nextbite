package models

import "time"

type Geo struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Store struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	PrimaryCuisine string    `json:"primary_cuisine"`
	Cuisines       []string  `json:"cuisines"`
	PriceTier      int       `json:"price_tier"`
	RatingAvg      float64   `json:"rating_avg"`
	RatingCount    int       `json:"rating_count"`
	Orders7D       int       `json:"orders_7d"`
	IsOpenNow      bool      `json:"is_open_now"`
	CreatedAt      time.Time `json:"created_at"`
	Geo            Geo       `json:"geo"`
}
