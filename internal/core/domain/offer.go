package domain

import "time"

type Offer struct {
	Title      string
	Price      float64
	Link       string
	Source     string
	ImageURL   string
	Tags       []string
	IsFeatured bool
	PostDate   time.Time
}
